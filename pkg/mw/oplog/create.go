package oplog

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/oplog/mw/v1/oplog"
	oplogcrud "github.com/NpoolPlatform/oplog-middleware/pkg/crud/oplog"
	"github.com/NpoolPlatform/oplog-middleware/pkg/db"
	"github.com/NpoolPlatform/oplog-middleware/pkg/db/ent"
)

func (h *Handler) CreateOpLog(ctx context.Context) (*npool.OpLog, error) {
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := oplogcrud.CreateSet(
			cli.OpLog.Create(),
			&oplogcrud.Req{
				AppID:         h.AppID,
				UserID:        h.UserID,
				Path:          h.Path,
				Method:        h.Method,
				Arguments:     h.Arguments,
				CurValue:      h.CurValue,
				HumanReadable: h.HumanReadable,
				Result:        h.Result,
				FailReason:    h.FailReason,
			},
		).Save(_ctx)
		if err != nil {
			return err
		}
		h.AutoID = &info.AutoID
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetOpLog(ctx)
}

func (h *Handler) CreateOpLogs(ctx context.Context) ([]*npool.OpLog, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}
