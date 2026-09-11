package service

import (
	"context"
	"fmt"
	"graph/param"
)

func (s Service) GetTasks(ctx context.Context, req param.GetTasksRequest) (param.GetAllTasksResponse, error){

	    if req.Page < 1 {
        req.Page = 1
    }

    if req.PageSize < 1 {
        req.PageSize = 10
    }

    if req.PageSize > 100 {
        req.PageSize = 100
    }

	tasks, total,  err := s.repo.GetTasks(ctx, req)
	if err != nil {
		return param.GetAllTasksResponse{}, fmt.Errorf("unexpected error:%w", err)
	}

	totalPages := (total + req.PageSize - 1) / req.PageSize

	    return param.GetAllTasksResponse{
        Tasks:      tasks,
        Page:       req.Page,
        PageSize:   req.PageSize,
        Total:      total,
        TotalPages: totalPages,
    }, nil
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