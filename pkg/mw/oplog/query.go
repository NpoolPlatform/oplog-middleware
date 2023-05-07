package oplog

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/oplog/mw/v1/oplog"
)

func (h *Handler) GetOpLog(ctx context.Context) (*npool.OpLog, error) {
	if h.AutoID == nil {
		return nil, fmt.Errorf("invalid auto_id")
	}
	return nil, nil
}

func (h *Handler) GetOpLogs(ctx context.Context) ([]*npool.OpLog, uint32, error) {
	return nil, 0, nil
}
