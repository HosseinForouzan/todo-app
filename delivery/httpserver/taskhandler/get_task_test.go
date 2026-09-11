package taskhandler

import (
	"context"
	"errors"
	"graph/entity"
	"graph/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandler_GetTaskByID_Success(t *testing.T) {
	mockRepo := &mockRepository{
		getTaskByIDFunc: func(ctx context.Context, id uint) (entity.Task, error) {
			return entity.Task{
				ID:          id,
				Title:       "Learn Go",
				Description: "Testing handlers",
				Status:      entity.StatusTodo,
				Assignee:    "Hossein",
			}, nil
		},
	}

	svc := service.New(mockRepo, nil)
	handler := New(svc)

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.GET("/tasks/:id", handler.GetTaskByID)

	req := httptest.NewRequest(
		http.MethodGet,
		"/tasks/1",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			rec.Code,
		)
	}

	body := rec.Body.String()

	if !strings.Contains(body, `"id":1`) {
		t.Errorf(
			"expected response to contain id 1, got %s",
			body,
		)
	}

	if !strings.Contains(body, `"title":"Learn Go"`) {
		t.Errorf(
			"expected response to contain title, got %s",
			body,
		)
	}
}

func TestHandler_GetTaskByID_InvalidID(t *testing.T) {
	mockRepo := &mockRepository{}

	svc := service.New(mockRepo, nil)
	handler := New(svc)

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.GET("/tasks/:id", handler.GetTaskByID)

	req := httptest.NewRequest(
		http.MethodGet,
		"/tasks/abc",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}
}

func TestHandler_GetTaskByID_ServiceError(t *testing.T) {
	repositoryErr := errors.New("database error")

	mockRepo := &mockRepository{
		getTaskByIDFunc: func(
			ctx context.Context,
			id uint,
		) (entity.Task, error) {
			return entity.Task{}, repositoryErr
		},
	}

	svc := service.New(mockRepo, nil)
	handler := New(svc)

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.GET("/tasks/:id", handler.GetTaskByID)

	req := httptest.NewRequest(
		http.MethodGet,
		"/tasks/1",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}

	if !strings.Contains(rec.Body.String(), "database error") {
		t.Errorf(
			"expected database error in response, got %s",
			rec.Body.String(),
		)
	}
}


func TestHandler_GetTasks_Success(t *testing.T) {
	mockRepo := &mockRepository{
		getTasksFunc: func(
			ctx context.Context,
		) ([]entity.Task, error) {
			return []entity.Task{
				{
					ID:          1,
					Title:       "Task 1",
					Description: "First task",
					Status:      entity.StatusTodo,
					Assignee:    "Hossein",
				},
				{
					ID:          2,
					Title:       "Task 2",
					Description: "Second task",
					Status:      entity.StatusDone,
					Assignee:    "Ali",
				},
			}, nil
		},
	}

	svc := service.New(mockRepo, nil)
	handler := New(svc)

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.GET("/tasks", handler.GetTasks)

	req := httptest.NewRequest(
		http.MethodGet,
		"/tasks",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			rec.Code,
		)
	}

	body := rec.Body.String()

	if !strings.Contains(body, `"id":1`) {
		t.Errorf("expected task 1 in response, got %s", body)
	}

	if !strings.Contains(body, `"id":2`) {
		t.Errorf("expected task 2 in response, got %s", body)
	}

	if !strings.Contains(body, `"title":"Task 1"`) {
		t.Errorf("expected Task 1 in response, got %s", body)
	}
}

func TestHandler_GetTasks_ServiceError(t *testing.T) {
	repositoryErr := errors.New("database error")

	mockRepo := &mockRepository{
		getTasksFunc: func(
			ctx context.Context,
		) ([]entity.Task, error) {
			return nil, repositoryErr
		},
	}

	svc := service.New(mockRepo, nil)
	handler := New(svc)

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.GET("/tasks", handler.GetTasks)

	req := httptest.NewRequest(
		http.MethodGet,
		"/tasks",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}

	if !strings.Contains(rec.Body.String(), "database error") {
		t.Errorf(
			"expected database error in response, got %s",
			rec.Body.String(),
		)
	}
}