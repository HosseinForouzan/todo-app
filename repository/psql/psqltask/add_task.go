package psqltask

import (
	"context"
	"fmt"
	"graph/entity"
)

func (d *DB) AddTask(ctx context.Context, task entity.Task) (entity.Task, error) {
	var id uint
	query := `INSERT INTO tasks(title,description,assignee)
	 VALUES($1,$2,$3) RETURNING id`

	err := d.conn.Pool().QueryRow(ctx, query, task.Title, task.Description, task.Assignee).Scan(&id)
	if err != nil {
		return entity.Task{}, fmt.Errorf("can't insert task: %w", err)
	}

	task.ID = id

	return task, nil
}
