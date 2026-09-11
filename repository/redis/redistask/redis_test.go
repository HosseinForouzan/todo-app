package redistask

import (
	"context"
	"strconv"
	"strings"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"graph/entity"
)

func setupAdapter(t *testing.T) Adapter {
	t.Helper()

	mr := miniredis.RunT(t)

	parts := strings.Split(mr.Addr(), ":")
	port, _ := strconv.Atoi(parts[1])

	return New(Config{
		Host:     parts[0],
		Port:     port,
		Password: "",
		DB:       0,
	})
}

func TestRedis_SetAndGetTask(t *testing.T) {
	ctx := context.Background()
	adapter := setupAdapter(t)

	task := entity.Task{
		ID:       1,
		Title:    "Test Task",
		Assignee: "Hossein",
	}

	err := adapter.SetTask(ctx, task)
	if err != nil {
		t.Fatalf("SetTask() error = %v", err)
	}

	got, err := adapter.GetTask(ctx, task.ID)
	if err != nil {
		t.Fatalf("GetTask() error = %v", err)
	}

	if got.ID != task.ID {
		t.Errorf("expected ID %d, got %d", task.ID, got.ID)
	}

	if got.Title != task.Title {
		t.Errorf("expected title %q, got %q", task.Title, got.Title)
	}

	if got.Assignee != task.Assignee {
		t.Errorf("expected assignee %q, got %q", task.Assignee, got.Assignee)
	}
}

func TestRedis_GetTask_NotFound(t *testing.T) {
	ctx := context.Background()
	adapter := setupAdapter(t)

	_, err := adapter.GetTask(ctx, 999)
	if err == nil {
		t.Fatal("expected error for non-existent key, got nil")
	}
}

func TestRedis_DeleteTask(t *testing.T) {
	ctx := context.Background()
	adapter := setupAdapter(t)

	task := entity.Task{
		ID:    2,
		Title: "To be deleted",
	}

	err := adapter.SetTask(ctx, task)
	if err != nil {
		t.Fatalf("SetTask() error = %v", err)
	}

	err = adapter.DeleteTask(ctx, task.ID)
	if err != nil {
		t.Fatalf("DeleteTask() error = %v", err)
	}

	_, err = adapter.GetTask(ctx, task.ID)
	if err == nil {
		t.Fatal("expected error after deletion, got nil")
	}
}

func TestRedis_DeleteTask_NonExistent(t *testing.T) {
	ctx := context.Background()
	adapter := setupAdapter(t)

	err := adapter.DeleteTask(ctx, 999)
	if err != nil {
		t.Fatalf("DeleteTask() on non-existent key should not error, got %v", err)
	}
}