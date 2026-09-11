package httpserver

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"graph/entity"
	"graph/param"
	"graph/service"

	"github.com/gin-gonic/gin"
)

type mockRepo struct{}

func (m *mockRepo) AddTask(ctx context.Context, task entity.Task) (entity.Task, error) {
	return task, nil
}
func (m *mockRepo) GetTaskByID(ctx context.Context, id uint) (entity.Task, error) {
	return entity.Task{ID: id, Title: "Test"}, nil
}
func (m *mockRepo) GetTasks(ctx context.Context, req param.GetTasksRequest) ([]entity.Task, int, error) {
	return []entity.Task{}, 0, nil
}
func (m *mockRepo) UpdateTask(ctx context.Context, task entity.Task) (entity.Task, error) {
	return task, nil
}
func (m *mockRepo) DeleteTask(ctx context.Context, id uint) error { return nil }
func (m *mockRepo) DoesTaskExist(ctx context.Context, id uint) (bool, error) {
	return true, nil
}

func setupTestServer(t *testing.T) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)

	svc := service.New(&mockRepo{}, nil)
	server := New(svc)

	server.Router.GET("/health-check", server.Healthcheck)
	server.Router.GET("/metrics", func(c *gin.Context) { c.Status(http.StatusOK) })
	server.Handler.SetRoutes(server.Router)

	return server.Router
}

func TestServer_New(t *testing.T) {
	gin.SetMode(gin.TestMode)
	svc := service.New(&mockRepo{}, nil)
	server := New(svc)

	if server.Router == nil {
		t.Fatal("expected Router to be initialized")
	}
}

func TestServer_HealthCheck(t *testing.T) {
	router := setupTestServer(t)

	req := httptest.NewRequest(http.MethodGet, "/health-check", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "ok") {
		t.Errorf("expected 'ok' in response body, got %s", rec.Body.String())
	}
}

func TestServer_MetricsEndpoint(t *testing.T) {
	router := setupTestServer(t)

	req := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestServer_TaskRoutes_Registered(t *testing.T) {
	router := setupTestServer(t)

	tests := []struct {
		method string
		path   string
	}{
		{http.MethodGet, "/task/"},
		{http.MethodPost, "/task/"},
		{http.MethodGet, "/task/1"},
		{http.MethodPut, "/task/1"},
		{http.MethodDelete, "/task/1"},
	}

	for _, tt := range tests {
		t.Run(tt.method+" "+tt.path, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if rec.Code == http.StatusNotFound {
				t.Errorf("route %s %s is not registered", tt.method, tt.path)
			}
		})
	}
}

func TestServer_Swagger_Registered(t *testing.T) {
	router := setupTestServer(t)

	req := httptest.NewRequest(http.MethodGet, "/swagger/index.html", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code == http.StatusNotFound {
		t.Error("swagger route is not registered")
	}
}

func TestServer_InitTracer(t *testing.T) {
	ctx := context.Background()

	shutdown, err := InitTracer(ctx)
	if err != nil {
		t.Fatalf("InitTracer() error = %v", err)
	}

	if shutdown == nil {
		t.Fatal("expected shutdown function, got nil")
	}

	if err := shutdown(ctx); err != nil {
		t.Errorf("shutdown() error = %v", err)
	}
}