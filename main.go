package main

import (
	"context"
	"fmt"
	"graph/param"
	"graph/repository/psql"
	"graph/repository/psql/psqltask"
	"graph/service"
)

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

	t, err := taskSvc.AddTask(ctx, param.AddTaskRequest{
		Title: "salam",
		Description: "HAJI",
		Assignee: "Hossein",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t)

	taskByID, err := taskSvc.GetTaskByID(ctx, param.GetTaskRequest{ID: 1})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(taskByID)
}