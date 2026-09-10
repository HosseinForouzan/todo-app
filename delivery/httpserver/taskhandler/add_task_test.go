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

func TestHandler_AddTask_Success(t *testing.T) {
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

	svc := service.New(mockRepo)
	handler := New(svc)

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.POST("/tasks", handler.AddTask)

	body := `{
		"title": "Learn Handler Testing",
		"description": "Write HTTP tests",
		"assignee": "Hossein"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/tasks",
		strings.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusCreated,
			rec.Code,
		)
	}

	if !strings.Contains(rec.Body.String(), `"id":1`) {
		t.Errorf(
			"expected response to contain task ID, got %s",
			rec.Body.String(),
		)
	}

	if !strings.Contains(rec.Body.String(), `"title":"Learn Handler Testing"`) {
		t.Errorf(
			"expected response to contain task title, got %s",
			rec.Body.String(),
		)
	}
}

func TestHandler_AddTask_InvalidJSON(t *testing.T) {
	mockRepo := &mockRepository{}

	svc := service.New(mockRepo)
	handler := New(svc)

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.POST("/tasks", handler.AddTask)

	body := `{
		"title": "Invalid JSON"
		"description": "missing comma"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/tasks",
		strings.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

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

func TestHandler_AddTask_ServiceError(t *testing.T) {
	repositoryErr := errors.New("database error")

	mockRepo := &mockRepository{
		addTaskFunc: func(ctx context.Context,task entity.Task) (entity.Task, error) {
			return entity.Task{}, repositoryErr
		},
	}

	svc := service.New(mockRepo)
	handler := New(svc)

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.POST("/tasks", handler.AddTask)

	body := `{
		"title": "Test Error",
		"description": "Repository error",
		"assignee": "Hossein"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/tasks",
		strings.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

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
			"expected error message in response, got %s",
			rec.Body.String(),
		)
	}
}
