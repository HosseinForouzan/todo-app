package param

import (
	"graph/entity"
	"time"
)

type GetTaskRequest struct {
	ID uint `json:"id"`
}

type GetTaskResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Assignee    string `json:"assignee"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetAllTasksResponse struct {
	Tasks []entity.Task
}