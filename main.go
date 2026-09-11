package main

import (
	"context"
	"database/sql"
	"fmt"
	"graph/delivery/httpserver"
	"graph/metrics"
	"graph/param"
	"graph/repository/psql"
	"graph/repository/psql/psqltask"
	"graph/repository/redis/redistask"
	"graph/service"
	"log"
	"os"

	migrate "github.com/rubenv/sql-migrate"
    _ "github.com/jackc/pgx/v5/stdlib"

	_ "graph/docs"

	"net/http"
	_ "net/http/pprof"
)

// @title           Task Manager API
// @version         1.0
// @description     A simple task management microservice
// @host            localhost:8080
// @BasePath        /
func main() {

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbPort := 5431
	if os.Getenv("DB_PORT") == "5432" {
		dbPort = 5432
	}

	config := psql.Config{
		Username: "myuser",
		Password: "secret",
		Host:     dbHost,
		Port:     dbPort,
		DBName:   "task_db",
	}
	ctx := context.Background()

	//migration
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Username, config.Password, config.Host, config.Port, config.DBName)

	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to open db for migration: %v", err)
	}
	migrations := &migrate.FileMigrationSource{Dir: "repository/psql/migrations"}
	n, err := migrate.Exec(sqlDB, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}
	fmt.Printf("applied %d migrations\n", n)
	sqlDB.Close()

	//Tracer
	shutdownTracer, err := httpserver.InitTracer(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer shutdownTracer(ctx)

	// Postgres and Redis
	psqlRepo, err := psql.NewPgxPool(ctx, config)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	psqltaskRepo := psqltask.New(psqlRepo)

	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost"
	}
	redisPort := 6380
	if os.Getenv("REDIS_PORT") == "6379" {
		redisPort = 6379
	}

	redsAdapter := redistask.New(redistask.Config{
		Host:     redisHost,
		Port:     redisPort,
		Password: "",
		DB:       0,
	})

	taskSvc := service.New(psqltaskRepo, redsAdapter)

	if resp, err := taskSvc.GetTasks(ctx, param.GetTasksRequest{Page: 1, PageSize: 10}); err == nil {
		metrics.TasksCount.Set(float64(len(resp.Tasks)))
	}

	go func() {
		fmt.Println("pprof listening on :6060")
		if err := http.ListenAndServe(":6060", nil); err != nil {
			fmt.Println("pprof error:", err)
		}
	}()

	server := httpserver.New(taskSvc)

	server.Serve()

}
