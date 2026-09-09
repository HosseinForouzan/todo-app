package httpserver

import (
	"graph/delivery/httpserver/taskhandler"
	"graph/service"
	"log"

	"github.com/gin-gonic/gin"
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
	s.Router.GET("/health-check", s.Healthcheck)
	s.Handler.SetRoutes(s.Router)

	if err := s.Router.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)

	}
}
