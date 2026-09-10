package httpserver

import (
	"graph/delivery/httpserver/middleware"
	"graph/delivery/httpserver/taskhandler"
	"graph/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type Server struct {
	Handler taskhandler.Handler
	Router  *gin.Engine
}

func New(taskSvc service.Service) Server {
	return Server{
		Handler: taskhandler.New(taskSvc),
		Router:  gin.Default(),
	}
}

func (s Server) Serve() {
	s.Router.Use(otelgin.Middleware("graph"))
	s.Router.Use(middleware.Metrics())

	s.Router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	s.Router.GET("/health-check", s.Healthcheck)
	s.Handler.SetRoutes(s.Router)

	if err := s.Router.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)

	}
}
