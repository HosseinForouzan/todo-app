package param

import "time"

type UpdateTaskRequest struct {
	ID uint `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Assignee    string `json:"assignee"`
}

type UpdateTaskResponse struct {
	ID uint `json:"id"`
	Title     string `json:"title"`
	UpdatedAt time.Time `json:"updated_at"`
}