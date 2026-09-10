package main

import (
	"context"
	"fmt"
	"graph/delivery/httpserver"
	"graph/repository/psql"
	"graph/repository/psql/psqltask"
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
		Host: "localhost",
		Port: 5431,
		DBName: "task_db",
		
	}
	ctx := context.Background()
	psqlRepo, err := psql.NewPgxPool(ctx, config)
	if err != nil {
		fmt.Println(err)
	}
	psqltaskRepo := psqltask.New(psqlRepo)

	taskSvc := service.New(psqltaskRepo)

	fmt.Println(taskSvc)

	server := httpserver.New(taskSvc)

	server.Serve()

}