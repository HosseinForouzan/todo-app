package service

import (
	"context"
	"errors"
	"testing"

	"graph/entity"
	"graph/param"
)

func TestService_AddTask_Success(t *testing.T) {
	mockRepo := &mockRepository{
		addTaskFunc: func(ctx context.Context, task entity.Task) (entity.Task, error) {
			return entity.Task{
				ID:          1,
				Title:       task.Title,
				Description: task.Description,
				Assignee:    task.Assignee,
			}, nil
		},
	}

	svc := New(mockRepo)

	req := param.AddTaskRequest{
		Title:       "Graph Task",
		Description: "Write unit tests",
		Assignee:    "Hossein",
	}

	got, err := svc.AddTask(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got.ID != 1 {
		t.Errorf("expected ID 1, got %d", got.ID)
	}

	if got.Title != req.Title {
		t.Errorf("expected title %q, got %q", req.Title, got.Title)
	}
}

func TestService_AddTask_RepositoryError(t *testing.T) {
	repositoryErr := errors.New("database error")

	mockRepo := &mockRepository{
		addTaskFunc: func(ctx context.Context, task entity.Task) (entity.Task, error) {
			return entity.Task{}, repositoryErr
		},
	}

	svc := New(mockRepo)

	req := param.AddTaskRequest{
		Title:       "Graph Task",
		Description: "Write unit tests",
		Assignee:    "Hossein",
	}

	_, err := svc.AddTask(context.Background(), req)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, repositoryErr) {
		t.Errorf("expected wrapped repository error, got %v", err)
	}
}
