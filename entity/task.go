
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
    ID          uint
    Title       string
    Description string
    Status      TaskStatus
    Assignee    string
    CreatedAt   time.Time
    UpdatedAt   time.Time
}