package oplog

import (
	"context"

	"github.com/NpoolPlatform/oplog-middleware/pkg/db"
	"github.com/NpoolPlatform/oplog-middleware/pkg/db/ent"

	oplogcrud "github.com/NpoolPlatform/oplog-middleware/pkg/crud/oplog"
)

func (h *Handler) CountOpLogs(ctx context.Context) (uint32, error) {
	count := uint32(0)

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := oplogcrud.SetQueryConds(cli.OpLog.Query(), h.Conds)
		if err != nil {
			return err
		}
		_count, err := stm.Count(_ctx)
		if err != nil {
			return err
		}
		count = uint32(_count)
		return nil
	})
	if err != nil {
		return 0, err
	}

	return count, nil
}
