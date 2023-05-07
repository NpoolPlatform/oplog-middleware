package oplog

import (
	"context"
)

func (h *Handler) ExistOpLog(ctx context.Context) (bool, error) {
	return false, nil
}

func (h *Handler) ExistOpLogConds(ctx context.Context) (bool, error) {
	return false, nil
}
