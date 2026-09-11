package psqltask

import (
	"context"
	"testing"

	"graph/entity"
	"graph/param"
	"graph/repository/psql"
)

func TestDB_GetTaskByID(t *testing.T) {
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
	t.Cleanup(func() { conn.Close() })

	repo := New(conn)

	createdTask, err := repo.AddTask(ctx, entity.Task{
		Title:       "GetTaskByID Test",
		Description: "Testing get task by id",
		Assignee:    "Hossein",
	})
	if err != nil {
		t.Fatalf("AddTask() error = %v", err)
	}

	got, err := repo.GetTaskByID(ctx, createdTask.ID)
	if err != nil {
		t.Fatalf("GetTaskByID() error = %v", err)
	}

	t.Cleanup(func() {
		repo.DeleteTask(ctx, got.ID)
	})

	if got.ID != createdTask.ID {
		t.Errorf(
			"expected ID %d, got %d",
			createdTask.ID,
			got.ID,
		)
	}

	if got.Title != createdTask.Title {
		t.Errorf(
			"expected title %q, got %q",
			createdTask.Title,
			got.Title,
		)
	}

	if got.Description != createdTask.Description {
		t.Errorf(
			"expected description %q, got %q",
			createdTask.Description,
			got.Description,
		)
	}

	if got.Assignee != createdTask.Assignee {
		t.Errorf(
			"expected assignee %q, got %q",
			createdTask.Assignee,
			got.Assignee,
		)
	}
}

func TestDB_GetTasks(t *testing.T) {
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
	t.Cleanup(func() { conn.Close() })

	repo := New(conn)

	tasks := []entity.Task{
		{Title: "Integration Test 1", Description: "First task", Assignee: "Hossein"},
		{Title: "Integration Test 2", Description: "Second task", Assignee: "Ali"},
		{Title: "Integration Test 3", Description: "Third task", Assignee: "Reza"},
	}

	for _, task := range tasks {
		createdTask, err := repo.AddTask(ctx, task)
		if err != nil {
			t.Fatalf("AddTask() error = %v", err)
		}
		t.Cleanup(func() {
			repo.DeleteTask(ctx, createdTask.ID)
		})
	}

	req := param.GetTasksRequest{Page: 1, PageSize: 10}

	got, total, err := repo.GetTasks(ctx, req)
	if err != nil {
		t.Fatalf("GetTasks() error = %v", err)
	}

	if len(got) < len(tasks) {
		t.Fatalf("expected at least %d tasks, got %d", len(tasks), len(got))
	}

	if total < len(tasks) {
		t.Fatalf("expected total at least %d, got %d", len(tasks), total)
	}
}
