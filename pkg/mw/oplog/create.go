package oplog

import (
	"context"

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
				AppID:     h.AppID,
				UserID:    h.UserID,
				Path:      h.Path,
				Method:    h.Method,
				Arguments: h.Arguments,
				CurValue:  h.CurValue,
			},
		).Save(_ctx)
		if err != nil {
			return err
		}
		h.EntID = &info.EntID
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetOpLog(ctx)
}
