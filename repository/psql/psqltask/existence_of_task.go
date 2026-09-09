package psqltask

import (
	"context"
	"database/sql"
	"fmt"
)

func (d *DB) DoesTaskExist(ctx context.Context, id uint) (bool, error) {
	var i uint
	query := `SELECT id FROM tasks WHERE id = $1`

	err := d.conn.Pool().QueryRow(ctx, query, id).Scan(&i)
	fmt.Println(sql.ErrNoRows)
	if err != nil {
		if err == sql.ErrNoRows{
			return false, fmt.Errorf("task doesn't exist")
		}

		return false, fmt.Errorf("can't get task: %w", err)
	}

	return true, nil
}
