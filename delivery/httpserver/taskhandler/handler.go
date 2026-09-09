package taskhandler

import "graph/service"

type Hadler struct {
	taskSvc service.Service
}

func New(taskSvc service.Service) Hadler {
	return Hadler{taskSvc: taskSvc}
}