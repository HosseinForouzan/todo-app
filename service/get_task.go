package service

import (
	"context"
	"fmt"
	"graph/param"
)

func (s Service) GetTasks(ctx context.Context) (param.GetAllTasksResponse, error){
	tasks, err := s.repo.GetTasks(ctx)
	if err != nil {
		return param.GetAllTasksResponse{}, fmt.Errorf("unexpected error:%w", err)
	}

	return param.GetAllTasksResponse{
		Tasks: tasks,
	},nil
}

func (s Service) GetTaskByID(ctx context.Context, req param.GetTaskRequest) (param.GetTaskResponse, error) {
	task, err := s.cache.GetTask(ctx, req.ID)
	if err == nil {
		return param.GetTaskResponse{
		ID: task.ID,
		Title: task.Title,
		Description: task.Description,
		Status: string(task.Status),
		Assignee: task.Assignee,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}, nil
	}

	task, err = s.repo.GetTaskByID(ctx, req.ID)
	if err != nil {
		return param.GetTaskResponse{}, fmt.Errorf("unexpected error:%w", err)
	}

	err = s.cache.SetTask(ctx, task)
	if err != nil {
		return param.GetTaskResponse{}, fmt.Errorf("unexpected error:%w", err)
	}

	return param.GetTaskResponse{
		ID: task.ID,
		Title: task.Title,
		Description: task.Description,
		Status: string(task.Status),
		Assignee: task.Assignee,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}, nil
}