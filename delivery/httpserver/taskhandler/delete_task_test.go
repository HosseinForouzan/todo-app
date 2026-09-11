package taskhandler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"graph/service"

	"github.com/gin-gonic/gin"
)

func TestHandler_DeleteTask_Success(t *testing.T) {
	var capturedID uint

	mockRepo := &mockRepository{
		doesTaskExistFunc: func(
			ctx context.Context,
			id uint,
		) (bool, error) {
			capturedID = id
			return true, nil
		},
		deleteTaskFunc: func(
			ctx context.Context,
			id uint,
		) error {
			capturedID = id
			return nil
		},
	}

	svc := service.New(mockRepo, nil)
	handler := New(svc)

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.DELETE("/tasks/:id", handler.DeleteTask)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/tasks/1",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			rec.Code,
		)
	}

	if capturedID != 1 {
		t.Errorf(
			"expected ID 1, got %d",
			capturedID,
		)
	}
}

func TestHandler_DeleteTask_InvalidID(t *testing.T) {
	mockRepo := &mockRepository{}

	svc := service.New(mockRepo, nil)
	handler := New(svc)

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.DELETE("/tasks/:id", handler.DeleteTask)

	req := httptest.NewRequest(
		http.MethodDelete,
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

func TestHandler_DeleteTask_ServiceError(t *testing.T) {
	repositoryErr := errors.New("database error")

	mockRepo := &mockRepository{
		doesTaskExistFunc: func(
			ctx context.Context,
			id uint,
		) (bool, error) {
			return true, nil
		},
		deleteTaskFunc: func(
			ctx context.Context,
			id uint,
		) error {
			return repositoryErr
		},
	}

	svc := service.New(mockRepo, nil)
	handler := New(svc)

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.DELETE("/tasks/:id", handler.DeleteTask)

	req := httptest.NewRequest(
		http.MethodDelete,
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