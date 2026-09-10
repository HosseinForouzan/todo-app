package psqltask

import (
	"context"
	"testing"

	"graph/entity"
	"graph/repository/psql"
)

func TestDB_DoesTaskExist(t *testing.T) {
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
		Title:       "Existence Test",
		Description: "Testing task existence",
		Assignee:    "Hossein",
	})
	if err != nil {
		t.Fatalf("AddTask() error = %v", err)
	}



	exists, err := repo.DoesTaskExist(ctx, createdTask.ID)
	if err != nil {
		t.Fatalf("DoesTaskExist() error = %v", err)
	}

	if !exists {
		t.Fatalf(
			"expected task with ID %d to exist",
			createdTask.ID,
		)
	}

	exists, err = repo.DoesTaskExist(ctx, 999)
	if err == nil {
		t.Fatalf("DoesTaskExist() error = %v", err)
	}

	if exists {
		t.Fatal("expected non-existing task to return false")
	}
}