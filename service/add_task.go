package service

import (
	"context"
	"fmt"
	"graph/entity"
	"graph/param"
)

func (s Service) AddTask(ctx context.Context, req param.AddTaskRequest) (param.AddTaskResponse, error) {
	task := entity.Task{
		ID: 0,
		Title: req.Title,
		Description: req.Description,
		Assignee: req.Assignee,
	}

	createdTask, err := s.repo.AddTask(ctx, task)
	if err != nil {
		return param.AddTaskResponse{}, fmt.Errorf("unexpected error:%w", err)
	}

	return param.AddTaskResponse{
		ID: createdTask.ID,
		Title: createdTask.Title,
	}, nil
}