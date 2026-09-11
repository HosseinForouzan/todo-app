package service

import (
	"context"
	"fmt"
	"graph/metrics"
	"graph/param"
)

func (s Service) DeleteTask(ctx context.Context, req param.DeleteTaskRequest) error {
	if doesTaskExist, _ := s.repo.DoesTaskExist(ctx, req.ID); !doesTaskExist {
		return fmt.Errorf("this task doesn't exist")
	}

	err := s.repo.DeleteTask(ctx, req.ID)
	if err != nil {
		return fmt.Errorf("unexpected error:%w", err)
	}

	metrics.TasksCount.Dec()

	err = s.cache.DeleteTask(ctx, req.ID)
	if err != nil {
		fmt.Println("failed to invalidate cache:%w", err)
	}

	return nil
}