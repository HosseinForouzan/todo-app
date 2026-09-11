package main

import (
	"context"
	"fmt"
	"graph/delivery/httpserver"
	"graph/metrics"
	"graph/param"
	"graph/repository/psql"
	"graph/repository/psql/psqltask"
	"graph/repository/redis/redistask"
	"graph/service"

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
		Host:     "localhost",
		Port:     6380,
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
