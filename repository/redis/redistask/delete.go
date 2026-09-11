package redistask

import (
	"context"
	"fmt"
)

func (a Adapter) DeleteTask(ctx context.Context, id uint) error {
	key := fmt.Sprintf("task:%d", id)

	return a.client.Del(ctx, key).Err()
}
