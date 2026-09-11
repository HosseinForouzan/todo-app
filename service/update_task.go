package service

import (
	"context"
	"fmt"
	"graph/entity"
	"graph/param"
)

func (s Service) UpdateTask(ctx context.Context, req param.UpdateTaskRequest)(param.UpdateTaskResponse, error) {
	task := entity.Task {
		ID: req.ID,
		Title: req.Title,
		Description: req.Description,
		Status: entity.TaskStatus(req.Status),
		Assignee: req.Assignee,
	}

	updatedTask, err := s.repo.UpdateTask(ctx, task)
	if err != nil {
		return param.UpdateTaskResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	err = s.cache.DeleteTask(ctx, req.ID)
	if err != nil {
		fmt.Println("failed to invlidate cache: %w", err.Error())
	}

	return param.UpdateTaskResponse{
		ID: updatedTask.ID,
		Title: updatedTask.Title,
		Description: updatedTask.Description,
		Status: string(updatedTask.Status),
		Assignee: updatedTask.Assignee,
		UpdatedAt: updatedTask.UpdatedAt,
	}, nil
}