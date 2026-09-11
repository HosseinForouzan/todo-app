package psqltask

import (
	"context"
	"testing"

	"graph/entity"
	"graph/repository/psql"
)

func TestDB_AddTask(t *testing.T) {
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
	t.Cleanup(func() {conn.Close()})

	repo := New(conn)

	

	task := entity.Task{
		Title:       "Integration Test",
		Description: "Testing PostgreSQL repository",
		Assignee:    "Hossein",
	}

	got, err := repo.AddTask(ctx, task)
	if err != nil {
		t.Fatalf("AddTask() error = %v", err)
	}

	t.Cleanup(func() {
		repo.DeleteTask(ctx, got.ID)
	})

	if got.ID == 0 {
		t.Fatal("expected task ID to be generated")
	}

	if got.Title != task.Title {
		t.Errorf("expected title %q, got %q", task.Title, got.Title)
	}

	if got.Description != task.Description {
		t.Errorf(
			"expected description %q, got %q",
			task.Description,
			got.Description,
		)
	}

	if got.Assignee != task.Assignee {
		t.Errorf(
			"expected assignee %q, got %q",
			task.Assignee,
			got.Assignee,
		)
	}
}