package psqltask

import (
	"context"
	"fmt"
	"graph/entity"
)

func (d *DB) GetTaskByID(ctx context.Context, id uint) (entity.Task, error) {
	var task entity.Task

	query := `SELECT * FROM tasks WHERE id = $1`
	err := d.conn.Pool().QueryRow(ctx, query, id).
	Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Assignee, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return entity.Task{}, fmt.Errorf("can't get task by id:%w", err)
	}

	return task, nil
}
