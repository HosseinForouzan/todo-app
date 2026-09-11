package main

import (
	"context"
	"fmt"
	"graph/delivery/httpserver"
	"graph/metrics"
	"graph/repository/psql"
	"graph/repository/psql/psqltask"
	"graph/repository/redis/redistask"
	"graph/service"

	_ "graph/docs"
)

// @title           Task Manager API
// @version         1.0
// @description     A simple task management microservice
// @host            localhost:8080
// @BasePath        /
func main() {

	config := psql.Config{
		Username: "myuser",
		Password: "secret",
		Host:     "localhost",
		Port:     5431,
		DBName:   "task_db",
	}
	ctx := context.Background()
	shutdownTracer, err := httpserver.InitTracer(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer shutdownTracer(ctx)

	psqlRepo, err := psql.NewPgxPool(ctx, config)
	if err != nil {
		fmt.Println(err)
	}
	psqltaskRepo := psqltask.New(psqlRepo)

	redsAdapter := redistask.New(redistask.Config{
		Host: "localhost",
		Port: 6380,
		Password: "",
		DB: 0,
	})

	fmt.Println(redsAdapter.Client().Ping(ctx))

	taskSvc := service.New(psqltaskRepo, nil)
	

	if resp, err := taskSvc.GetTasks(ctx); err == nil {
		metrics.TasksCount.Set(float64(len(resp.Tasks)))
	}

	fmt.Println(taskSvc)

	server := httpserver.New(taskSvc)

	server.Serve()

}
