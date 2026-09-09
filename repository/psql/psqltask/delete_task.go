package psqltask

import (
	"context"
	"fmt"
)

func (d *DB) DeleteTask(ctx context.Context, id uint) error {
	query := `DELETE FROM tasks WHERE id = $1`
	_, err := d.conn.Pool().Exec(ctx, query, id)
	if err != nil {
		fmt.Errorf("can't delete task:%w", err)
	}

	return nil
}