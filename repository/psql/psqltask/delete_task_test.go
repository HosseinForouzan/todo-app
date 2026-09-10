package psqltask

import (
	"context"
	"testing"

	"graph/entity"
	"graph/repository/psql"
)

func TestDB_DeleteTask(t *testing.T) {
	ctx := context.Background()

	config := psql.Config{
		Username: "myuser",
		Password: "secret",
		Host:     "localhost",
		Port:     5431,
		DBName:   "task_db",
	}

	conn, err := psql.NewPgxPool(ctx, config)
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	defer conn.Close()

	repo := New(conn)

	createdTask, err := repo.AddTask(ctx, entity.Task{
		Title:       "Delete Test",
		Description: "This task should be deleted",
		Assignee:    "Hossein",
	})
	if err != nil {
		t.Fatalf("AddTask() error = %v", err)
	}

	err = repo.DeleteTask(ctx, createdTask.ID)
	if err != nil {
		t.Fatalf("DeleteTask() error = %v", err)
	}

	_, err = repo.GetTaskByID(ctx, createdTask.ID)
	if err == nil {
		t.Fatal("expected error when getting deleted task, got nil")
	}
}