package redistask

import (
	"context"
	"encoding/json"
	"fmt"
	"graph/entity"
)

func (a Adapter) GetTask(ctx context.Context, id uint) (entity.Task, error) {
	var task entity.Task

	key := fmt.Sprintf("task:%d", id)

	value, err := a.client.Get(ctx, key).Result()
	
	if err != nil {
		return task, err
	}

	err = json.Unmarshal([]byte(value), &task)
	if err != nil {
		return task, err
	}

	return task, nil
}
