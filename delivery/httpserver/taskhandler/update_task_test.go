package taskhandler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"graph/entity"
	"graph/service"

	"github.com/gin-gonic/gin"
)

func TestHandler_UpdateTask_Success(t *testing.T) {
	var capturedTask entity.Task

	mockRepo := &mockRepository{
		updateTaskFunc: func(
			ctx context.Context,
			task entity.Task,
		) (entity.Task, error) {
			capturedTask = task

			return entity.Task{
				ID:        task.ID,
				Title:     task.Title,
				UpdatedAt: time.Now(),
			}, nil
		},
	}

	svc := service.New(mockRepo)
	handler := New(svc)

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.PUT("/tasks/:id", handler.UpdateTask)

	body := `{
		"title": "Updated Task",
		"description": "Updated description",
		"status": "in_progress",
		"assignee": "Ali"
	}`

	req := httptest.NewRequest(
		http.MethodPut,
		"/tasks/1",
		strings.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			rec.Code,
		)
	}

	// بررسی می‌کنیم ID از URL درست وارد Service شده.
	if capturedTask.ID != 1 {
		t.Errorf(
			"expected ID 1, got %d",
			capturedTask.ID,
		)
	}

	if capturedTask.Title != "Updated Task" {
		t.Errorf(
			"expected title %q, got %q",
			"Updated Task",
			capturedTask.Title,
		)
	}

	if capturedTask.Description != "Updated description" {
		t.Errorf(
			"expected description %q, got %q",
			"Updated description",
			capturedTask.Description,
		)
	}

	if capturedTask.Status != entity.StatusInProgress {
		t.Errorf(
			"expected status %q, got %q",
			entity.StatusInProgress,
			capturedTask.Status,
		)
	}

	if capturedTask.Assignee != "Ali" {
		t.Errorf(
			"expected assignee %q, got %q",
			"Ali",
			capturedTask.Assignee,
		)
	}

	bodyResponse := rec.Body.String()

	if !strings.Contains(bodyResponse, `"id":1`) {
		t.Errorf(
			"expected response to contain id 1, got %s",
			bodyResponse,
		)
	}

	if !strings.Contains(bodyResponse, `"title":"Updated Task"`) {
		t.Errorf(
			"expected response to contain updated title, got %s",
			bodyResponse,
		)
	}
}

func TestHandler_UpdateTask_InvalidID(t *testing.T) {
	mockRepo := &mockRepository{}

	svc := service.New(mockRepo)
	handler := New(svc)

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.PUT("/tasks/:id", handler.UpdateTask)

	body := `{
		"title": "Updated Task",
		"description": "Updated description",
		"status": "in_progress",
		"assignee": "Ali"
	}`

	req := httptest.NewRequest(
		http.MethodPut,
		"/tasks/abc",
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

func TestHandler_UpdateTask_InvalidJSON(t *testing.T) {
	mockRepo := &mockRepository{}

	svc := service.New(mockRepo)
	handler := New(svc)

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.PUT("/tasks/:id", handler.UpdateTask)

	body := `{
		"title": "Updated Task"
		"description": "Invalid JSON"
	}`

	req := httptest.NewRequest(
		http.MethodPut,
		"/tasks/1",
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

func TestHandler_UpdateTask_ServiceError(t *testing.T) {
	repositoryErr := errors.New("database error")

	mockRepo := &mockRepository{
		updateTaskFunc: func(
			ctx context.Context,
			task entity.Task,
		) (entity.Task, error) {
			return entity.Task{}, repositoryErr
		},
	}

	svc := service.New(mockRepo)
	handler := New(svc)

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.PUT("/tasks/:id", handler.UpdateTask)

	body := `{
		"title": "Updated Task",
		"description": "Updated description",
		"status": "in_progress",
		"assignee": "Ali"
	}`

	req := httptest.NewRequest(
		http.MethodPut,
		"/tasks/1",
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
			"expected database error in response, got %s",
			rec.Body.String(),
		)
	}
}