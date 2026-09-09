package taskhandler

import "graph/service"

type Handler struct {
	taskSvc service.Service
}

func New(taskSvc service.Service) Handler {
	return Handler{taskSvc: taskSvc}
}