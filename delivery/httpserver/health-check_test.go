package httpserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHealthcheck(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	s := Server{Router: router}
	router.GET("/health-check", s.Healthcheck)

	req := httptest.NewRequest(http.MethodGet, "/health-check", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}