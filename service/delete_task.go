package service

import (
	"context"
	"fmt"
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

	return nil
}