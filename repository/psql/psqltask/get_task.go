package psqltask

import (
	"context"
	"fmt"
	"graph/entity"
	"graph/param"
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

func (d *DB) GetTasks(ctx context.Context, req param.GetTasksRequest) ([]entity.Task, int, error) {
	var tasks []entity.Task
	var total int

	offset := (req.Page - 1) * req.PageSize

	countQuery := `SELECT COUNT(*) FROM tasks
		WHERE ($1 = '' OR status::text = $1)
		AND   ($2 = '' OR assignee    = $2)`

	err := d.conn.Pool().QueryRow(ctx, countQuery, req.Status, req.Assignee).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("can't count tasks: %w", err)
	}

	query := `SELECT * FROM tasks
		WHERE ($1 = '' OR status::text = $1)
		AND   ($2 = '' OR assignee    = $2)
		LIMIT $3 OFFSET $4`

	rows, err := d.conn.Pool().Query(ctx, query, req.Status, req.Assignee, req.PageSize, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("can't get tasks: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var task entity.Task
		rows.Scan(&task.ID, &task.Title, &task.Description,
			&task.Status, &task.Assignee, &task.CreatedAt, &task.UpdatedAt)
		tasks = append(tasks, task)
	}

	return tasks, total, nil
}