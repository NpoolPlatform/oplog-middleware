package oplog

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/oplog/mw/v1/oplog"
	oplogcrud "github.com/NpoolPlatform/oplog-middleware/pkg/crud/oplog"
	"github.com/NpoolPlatform/oplog-middleware/pkg/db"
	"github.com/NpoolPlatform/oplog-middleware/pkg/db/ent"
	entoplog "github.com/NpoolPlatform/oplog-middleware/pkg/db/ent/oplog"
)

func (h *Handler) UpdateOpLog(ctx context.Context) (*npool.OpLog, error) {
	if h.EntID == nil {
		return nil, fmt.Errorf("invalid ent_id")
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			OpLog.
			Query().
			Where(
				entoplog.EntID(*h.EntID),
			).
			Only(_ctx)
		if err != nil {
			return nil
		}

		info, err = oplogcrud.UpdateSet(
			info.Update(),
			&oplogcrud.Req{
				NewValue:      h.NewValue,
				HumanReadable: h.HumanReadable,
				Result:        h.Result,
				FailReason:    h.FailReason,
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
