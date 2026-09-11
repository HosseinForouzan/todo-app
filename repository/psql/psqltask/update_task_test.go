package psqltask

import (
	"context"
	"testing"

	"graph/entity"
	"graph/repository/psql"
)

func TestDB_UpdateTask(t *testing.T) {
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

	createdTask, err := repo.AddTask(ctx, entity.Task{
		Title:       "Before Update",
		Description: "Old description",
		Assignee:    "Hossein",
	})
	if err != nil {
		t.Fatalf("AddTask() error = %v", err)
	}

	t.Cleanup(func() {
		repo.DeleteTask(ctx,createdTask.ID)
	})

	taskToUpdate := entity.Task{
		ID:          createdTask.ID,
		Title:       "After Update",
		Description: "New description",
		Status:      entity.StatusInProgress,
		Assignee:    "Ali",
	}

	updatedTask, err := repo.UpdateTask(ctx, taskToUpdate)
	if err != nil {
		t.Fatalf("UpdateTask() error = %v", err)
	}

	if updatedTask.ID != createdTask.ID {
		t.Errorf(
			"expected ID %d, got %d",
			createdTask.ID,
			updatedTask.ID,
		)
	}

	if updatedTask.Title != taskToUpdate.Title {
		t.Errorf(
			"expected title %q, got %q",
			taskToUpdate.Title,
			updatedTask.Title,
		)
	}


	got, err := repo.GetTaskByID(ctx, createdTask.ID)
	if err != nil {
		t.Fatalf("GetTaskByID() after update error = %v", err)
	}

	if got.Title != taskToUpdate.Title {
		t.Errorf(
			"expected title %q after update, got %q",
			taskToUpdate.Title,
			got.Title,
		)
	}

	if got.Description != taskToUpdate.Description {
		t.Errorf(
			"expected description %q after update, got %q",
			taskToUpdate.Description,
			got.Description,
		)
	}

	if got.Status != taskToUpdate.Status {
		t.Errorf(
			"expected status %q after update, got %q",
			taskToUpdate.Status,
			got.Status,
		)
	}

	if got.Assignee != taskToUpdate.Assignee {
		t.Errorf(
			"expected assignee %q after update, got %q",
			taskToUpdate.Assignee,
			got.Assignee,
		)
	}
}