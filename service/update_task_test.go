package service

import (
	"context"
	"errors"
	"testing"

	"graph/entity"
	"graph/param"
)

func TestService_UpdateTask_Success(t *testing.T) {
	req := param.UpdateTaskRequest{
		ID:          10,
		Title:       "Updated Task",
		Description: "Updated Description",
		Status:      "in_progress",
		Assignee:    "Hossein",
	}

	expectedTask := entity.Task{
		ID:          10,
		Title:       "Updated Task",
		Description: "Updated Description",
		Status:      entity.TaskStatus("in_progress"),
		Assignee:    "Hossein",
	}

	mockRepo := &mockRepository{
		updateTaskFunc: func(ctx context.Context, task entity.Task) (entity.Task, error) {

			if task.ID != req.ID {
				t.Errorf("expected ID %d, got %d", req.ID, task.ID)
			}

			if task.Title != req.Title {
				t.Errorf("expected title %q, got %q", req.Title, task.Title)
			}

			if task.Description != req.Description {
				t.Errorf(
					"expected description %q, got %q",
					req.Description,
					task.Description,
				)
			}

			if task.Status != entity.TaskStatus(req.Status) {
				t.Errorf(
					"expected status %q, got %q",
					req.Status,
					task.Status,
				)
			}

			if task.Assignee != req.Assignee {
				t.Errorf(
					"expected assignee %q, got %q",
					req.Assignee,
					task.Assignee,
				)
			}

			return expectedTask, nil
		},
	}

	svc := New(mockRepo)

	got, err := svc.UpdateTask(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got.ID != expectedTask.ID {
		t.Errorf("expected ID %d, got %d", expectedTask.ID, got.ID)
	}

	if got.Title != expectedTask.Title {
		t.Errorf(
			"expected title %q, got %q",
			expectedTask.Title,
			got.Title,
		)
	}

	if got.Description != expectedTask.Description {
		t.Errorf(
			"expected description %q, got %q",
			expectedTask.Description,
			got.Description,
		)
	}

	if got.Status != string(expectedTask.Status) {
		t.Errorf(
			"expected status %q, got %q",
			expectedTask.Status,
			got.Status,
		)
	}

	if got.Assignee != expectedTask.Assignee {
		t.Errorf(
			"expected assignee %q, got %q",
			expectedTask.Assignee,
			got.Assignee,
		)
	}
}

func TestService_UpdateTask_RepositoryError(t *testing.T) {
	repositoryErr := errors.New("database error")

	mockRepo := &mockRepository{
		updateTaskFunc: func(
			ctx context.Context,
			task entity.Task,
		) (entity.Task, error) {
			return entity.Task{}, repositoryErr
		},
	}

	svc := New(mockRepo)

	req := param.UpdateTaskRequest{
		ID:          10,
		Title:       "Updated Task",
		Description: "Updated Description",
		Status:      "in_progress",
		Assignee:    "Hossein",
	}

	_, err := svc.UpdateTask(context.Background(), req)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, repositoryErr) {
		t.Errorf("expected repository error, got %v", err)
	}
}
