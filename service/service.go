package service

import (
	"context"
	"graph/entity"
)

type Repository interface {
	AddTask(ctx context.Context, task entity.Task) (entity.Task, error)
	GetTaskByID(ctx context.Context, id uint) (entity.Task, error)
	GetTasks(ctx context.Context) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task entity.Task) (entity.Task, error)
	DeleteTask(ctx context.Context, id uint) error
	DoesTaskExist(ctx context.Context, id uint) (bool, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}