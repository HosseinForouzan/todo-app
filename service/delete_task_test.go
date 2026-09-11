package service

import (
	"context"
	"errors"
	"graph/param"
	"testing"
)

func TestService_DeleteTask_TaskNotFound(t *testing.T) {
	deleteCalled := false

	mockRepo := &mockRepository{
		doesTaskExistFunc: func(ctx context.Context, id uint) (bool, error) {
			return false, nil
		},
		deleteTaskFunc: func(ctx context.Context, id uint) error {
			deleteCalled = true
			return nil
		},
	}

	svc := New(mockRepo, nil)
	req := param.DeleteTaskRequest{ID: 10}

	err := svc.DeleteTask(context.Background(), req)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if deleteCalled {
		t.Fatal("DeleteTask should not be called when task does not exist")
	}
}

func TestService_DeleteTask_Success(t *testing.T) {
	deleteCalled := false

	mockRepo := &mockRepository{
		doesTaskExistFunc: func(ctx context.Context, id uint) (bool, error) {
			return true, nil
		},
		deleteTaskFunc: func(ctx context.Context, id uint) error {
			deleteCalled = true
			return nil
		},
	}

	svc := New(mockRepo, nil)
	req := param.DeleteTaskRequest{ID: 10}

	err := svc.DeleteTask(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !deleteCalled {
		t.Fatal("expected DeleteTask to be called")
	}
}

func TestService_DeleteTask_RepositoryError(t *testing.T) {
	repositoryErr := errors.New("database error")

	mockRepo := &mockRepository{
		doesTaskExistFunc: func(ctx context.Context, id uint) (bool, error) {
			return true, nil
		},
		deleteTaskFunc: func(ctx context.Context, id uint) error {
			return repositoryErr
		},
	}

	svc := New(mockRepo, nil)
	req := param.DeleteTaskRequest{ID: 10}

	err := svc.DeleteTask(context.Background(), req)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, repositoryErr) {
		t.Errorf("expected repository error, got %v", err)
	}
}


func TestService_DeleteTask_ExistenceCheckError(t *testing.T) {
	repositoryErr := errors.New("this task doesn't exist")

	deleteCalled := false

	mockRepo := &mockRepository{
		doesTaskExistFunc: func(ctx context.Context, id uint) (bool, error) {
			return false, repositoryErr
		},
		deleteTaskFunc: func(ctx context.Context, id uint) error {
			deleteCalled = true
			return nil
		},
	}

	svc := New(mockRepo, nil)
	req := param.DeleteTaskRequest{ID: 10}

	err := svc.DeleteTask(context.Background(), req)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	

	if deleteCalled {
		t.Fatal("DeleteTask should not be called when DoesTaskExist returns an error")
	}
}