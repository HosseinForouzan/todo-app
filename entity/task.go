
package entity

import (
    "time"
)

type TaskStatus string

const (
    StatusTodo       TaskStatus = "todo"
    StatusInProgress TaskStatus = "in_progress"
    StatusDone       TaskStatus = "done"
)

func (s TaskStatus) IsValid() bool {
    switch s {
    case StatusTodo, StatusInProgress, StatusDone:
        return true
    }
    return false
}

type Task struct {
    ID          uint `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      TaskStatus `json:"status"`
    Assignee    string `json:"assignee"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}