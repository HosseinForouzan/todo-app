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

type GetTasksRequest struct {
	Page     int
	PageSize int
	Status   string
	Assignee string
}

type GetAllTasksResponse struct {
	Tasks      []entity.Task `json:"tasks"`
	Page       int           `json:"page"`
	PageSize   int           `json:"page_size"`
	Total      int           `json:"total"`
	TotalPages int           `json:"total_pages"`
}

