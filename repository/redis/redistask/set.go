package redistask

import (
	"context"
	"encoding/json"
	"fmt"
	"graph/entity"
	"time"

)

func (a Adapter) SetTask(ctx context.Context, task entity.Task) error {
	key := fmt.Sprintf("task:%d", task.ID)

	data, err := json.Marshal(task)
	if err != nil {
		return err
	}

	return a.client.Set(
		ctx,
		key,
		data,
		5*time.Minute,
	).Err()
}
