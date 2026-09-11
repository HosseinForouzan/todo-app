package service

import (
	"context"
	"errors"
	"graph/entity"
)

type mockRepository struct {
	addTaskFunc       func(ctx context.Context, task entity.Task) (entity.Task, error)
	getTaskByIDFunc   func(ctx context.Context, id uint) (entity.Task, error)
	getTasksFunc      func(ctx context.Context) ([]entity.Task, error)
	updateTaskFunc    func(ctx context.Context, task entity.Task) (entity.Task, error)
	deleteTaskFunc    func(ctx context.Context, id uint) error
	doesTaskExistFunc func(ctx context.Context, id uint) (bool, error)
}

func (m *mockRepository) AddTask(ctx context.Context, task entity.Task) (entity.Task, error) {
	return m.addTaskFunc(ctx, task)
}

func (m *mockRepository) GetTaskByID(ctx context.Context, id uint) (entity.Task, error) {
	return m.getTaskByIDFunc(ctx, id)
}

func (m *mockRepository) GetTasks(ctx context.Context) ([]entity.Task, error) {
	return m.getTasksFunc(ctx)
}

func (m *mockRepository) UpdateTask(ctx context.Context, task entity.Task) (entity.Task, error) {
	return m.updateTaskFunc(ctx, task)
}

func (m *mockRepository) DeleteTask(ctx context.Context, id uint) error {
	return m.deleteTaskFunc(ctx, id)
}

func (m *mockRepository) DoesTaskExist(ctx context.Context, id uint) (bool, error) {
	return m.doesTaskExistFunc(ctx, id)
}

type MockCache struct {
	tasks map[uint]entity.Task
}

func NewMockCache() *MockCache {
	return &MockCache{
		tasks: make(map[uint]entity.Task),
	}
}

func (m *MockCache) GetTask(ctx context.Context, id uint) (entity.Task, error) {
	task, ok := m.tasks[id]
	if !ok {
		return entity.Task{}, errors.New("task not found in cache")
	}

	return task, nil
}

func (m *MockCache) SetTask(ctx context.Context, task entity.Task) error {
	m.tasks[task.ID] = task
	return nil
}

func (m *MockCache) DeleteTask(ctx context.Context, id uint) error {
	delete(m.tasks, id)
	return nil
}

