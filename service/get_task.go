package service

import (
	"context"
	"fmt"
	"graph/param"
)

func (s Service) GetTaskByID(ctx context.Context, req param.GetTaskRequest) (param.GetTaskResponse, error) {
	task, err := s.repo.GetTaskByID(ctx, req.ID)
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