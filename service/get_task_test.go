package service

import (
	"context"
	"errors"
	"testing"

	"graph/entity"
	"graph/param"
)

func TestService_GetTaskByID_Success(t *testing.T) {
	expectedTask := entity.Task{
		ID:          10,
		Title:       "Learn Go",
		Description: "Write unit tests",
		Assignee:    "Hossein",
	}

	mockRepo := &mockRepository{
		getTaskByIDFunc: func(ctx context.Context, id uint) (entity.Task, error) {
			return expectedTask, nil
		},
	}
	ctx := context.Background()

	cache := NewMockCache()

	err := cache.SetTask(ctx, expectedTask)
	if err != nil {
		t.Fatal(err)
	}
	svc := New(mockRepo, cache)
	req := param.GetTaskRequest{ID: 10}

	got, err := svc.GetTaskByID(context.Background(), req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got.ID != expectedTask.ID {
		t.Errorf("expected ID %d, got %d", expectedTask.ID, got.ID)
	}

	if got.Title != expectedTask.Title {
		t.Errorf("expected title %q, got %q", expectedTask.Title, got.Title)
	}

	if got.Description != expectedTask.Description {
		t.Errorf("expected description %q, got %q", expectedTask.Description, got.Description)
	}

	if got.Assignee != expectedTask.Assignee {
		t.Errorf("expected assignee %q, got %q", expectedTask.Assignee, got.Assignee)
	}
}

func TestService_GetTaskByID_RepositoryError(t *testing.T) {
	repositoryErr := errors.New("database error")

	mockRepo := &mockRepository{
		getTaskByIDFunc: func(ctx context.Context, id uint) (entity.Task, error) {
			return entity.Task{}, repositoryErr
		},
	}

	cache := NewMockCache()
	ctx := context.Background()

	err := cache.SetTask(ctx, entity.Task{})
	if err != nil {
		t.Fatal(err)
	}

	svc := New(mockRepo, cache)
	req := param.GetTaskRequest{ID: 10}

	_, err = svc.GetTaskByID(context.Background(), req)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, repositoryErr) {
		t.Errorf("expected repository error, got %v", err)
	}
}



func TestService_GetTaskByID_CacheMiss_Success(t *testing.T) {
	expectedTask := entity.Task{
		ID:          10,
		Title:       "Learn Go",
		Description: "Write unit tests",
		Assignee:    "Hossein",
	}

	mockRepo := &mockRepository{
		getTaskByIDFunc: func(ctx context.Context, id uint) (entity.Task, error) {
			return expectedTask, nil
		},
	}

	cache := NewMockCache()
	svc := New(mockRepo, cache)

	got, err := svc.GetTaskByID(context.Background(), param.GetTaskRequest{ID: 10})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got.ID != expectedTask.ID {
		t.Errorf("expected ID %d, got %d", expectedTask.ID, got.ID)
	}

	cachedTask, err := cache.GetTask(context.Background(), 10)
	if err != nil {
		t.Error("expected task to be cached after cache miss")
	}
	if cachedTask.ID != expectedTask.ID {
		t.Errorf("expected cached ID %d, got %d", expectedTask.ID, cachedTask.ID)
	}
}





func TestService_GetTasks(t *testing.T) {
	repositoryErr := errors.New("database error")

	tests := []struct {
		name       string
		inputReq   param.GetTasksRequest
		repoTasks  []entity.Task
		repoErr    error
		wantTasks  []entity.Task
		wantErr    bool
		wantPage   int
		wantSize   int
		wantTotal  int
	}{
		{
			name:     "success",
			inputReq: param.GetTasksRequest{Page: 1, PageSize: 10},
			repoTasks: []entity.Task{
				{ID: 1, Title: "Task 1"},
				{ID: 2, Title: "Task 2"},
			},
			wantTasks: []entity.Task{
				{ID: 1, Title: "Task 1"},
				{ID: 2, Title: "Task 2"},
			},
			wantPage:  1,
			wantSize:  10,
			wantTotal: 2,
		},
		{
			name:      "empty result",
			inputReq:  param.GetTasksRequest{Page: 1, PageSize: 10},
			repoTasks: []entity.Task{},
			wantTasks: []entity.Task{},
			wantPage:  1,
			wantSize:  10,
			wantTotal: 0,
		},
		{
			name:    "repository error",
			inputReq: param.GetTasksRequest{Page: 1, PageSize: 10},
			repoErr: repositoryErr,
			wantErr: true,
		},
		{
			name:     "page zero defaults to 1",
			inputReq: param.GetTasksRequest{Page: 0, PageSize: 10},
			repoTasks: []entity.Task{
				{ID: 1, Title: "Task 1"},
			},
			wantTasks: []entity.Task{
				{ID: 1, Title: "Task 1"},
			},
			wantPage:  1,
			wantSize:  10,
			wantTotal: 1,
		},
		{
			name:     "pageSize zero defaults to 10",
			inputReq: param.GetTasksRequest{Page: 1, PageSize: 0},
			repoTasks: []entity.Task{
				{ID: 1, Title: "Task 1"},
			},
			wantTasks: []entity.Task{
				{ID: 1, Title: "Task 1"},
			},
			wantPage:  1,
			wantSize:  10,
			wantTotal: 1,
		},
		{
			name:     "pageSize over 100 capped to 100",
			inputReq: param.GetTasksRequest{Page: 1, PageSize: 200},
			repoTasks: []entity.Task{
				{ID: 1, Title: "Task 1"},
			},
			wantTasks: []entity.Task{
				{ID: 1, Title: "Task 1"},
			},
			wantPage:  1,
			wantSize:  100,
			wantTotal: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mockRepository{
				getTasksFunc: func(ctx context.Context, req param.GetTasksRequest) ([]entity.Task, int, error) {
					return tt.repoTasks, len(tt.repoTasks), tt.repoErr
				},
			}

			svc := New(mockRepo, nil)

			got, err := svc.GetTasks(context.Background(), tt.inputReq)

			if (err != nil) != tt.wantErr {
				t.Fatalf("expected error: %v, got: %v", tt.wantErr, err)
			}

			if tt.wantErr {
				if !errors.Is(err, repositoryErr) {
					t.Errorf("expected repository error, got %v", err)
				}
				return
			}

			if got.Page != tt.wantPage {
				t.Errorf("expected page %d, got %d", tt.wantPage, got.Page)
			}

			if got.PageSize != tt.wantSize {
				t.Errorf("expected pageSize %d, got %d", tt.wantSize, got.PageSize)
			}

			if got.Total != tt.wantTotal {
				t.Errorf("expected total %d, got %d", tt.wantTotal, got.Total)
			}

			if len(got.Tasks) != len(tt.wantTasks) {
				t.Fatalf("expected %d tasks, got %d", len(tt.wantTasks), len(got.Tasks))
			}

			for i := range tt.wantTasks {
				if got.Tasks[i].ID != tt.wantTasks[i].ID {
					t.Errorf("task %d: expected ID %d, got %d", i, tt.wantTasks[i].ID, got.Tasks[i].ID)
				}

				if got.Tasks[i].Title != tt.wantTasks[i].Title {
					t.Errorf("task %d: expected title %q, got %q", i, tt.wantTasks[i].Title, got.Tasks[i].Title)
				}
			}
		})
	}
}