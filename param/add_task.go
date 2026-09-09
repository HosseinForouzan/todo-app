package param

type AddTaskRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Assignee string `json:"assignee"`
}

type AddTaskResponse struct {
	ID uint `json:"id"`
	Title string `json:"title"`
}