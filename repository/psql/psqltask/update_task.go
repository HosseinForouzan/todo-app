package psqltask

import (
	"context"
	"fmt"
	"graph/entity"
	"time"
)

func (d *DB) UpdateTask(ctx context.Context, task entity.Task) (entity.Task, error) {
	var updatedAt time.Time

	query := `UPDATE tasks SET title=$1, description=$2, status=$3, assignee=$4, updated_at=Now()
				WHERE id = $5 RETURNING updated_at`
	err := d.conn.Pool().QueryRow(ctx, query, task.Title,
		 task.Description, task.Status, task.Assignee, task.ID).Scan(&updatedAt)
	if err != nil {
		return entity.Task{}, fmt.Errorf("can't update task: %w", err)
	}

	task.UpdatedAt = updatedAt

	return task, nil

}
