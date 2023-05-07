package oplog

import (
	"context"
	"fmt"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/oplog/mw/v1/oplog"
	oplogcrud "github.com/NpoolPlatform/oplog-middleware/pkg/crud/oplog"

	"github.com/NpoolPlatform/oplog-middleware/pkg/db"
	"github.com/NpoolPlatform/oplog-middleware/pkg/db/ent"
)

func (h *Handler) GetOpLog(ctx context.Context) (*npool.OpLog, error) {
	if h.AutoID == nil {
		return nil, fmt.Errorf("invalid auto_id")
	}

	h.Offset = 0
	h.Limit = 2
	h.Conds = &oplogcrud.Conds{
		AutoID: &cruder.Cond{Op: cruder.EQ, Val: *h.AutoID},
	}

	infos, _, err := h.GetOpLogs(ctx)
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}
	if len(infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	return infos[0], nil
}

func (h *Handler) GetOpLogs(ctx context.Context) ([]*npool.OpLog, uint32, error) {
	infos := []*npool.OpLog{}
	total := uint32(0)

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := oplogcrud.SetQueryConds(
			cli.OpLog.Query(),
			h.Conds,
		)
		if err != nil {
			return err
		}

		_total, err := stm.Count(_ctx)
		if err != nil {
			return err
		}

		total = uint32(_total)
		return stm.
			Select().
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, 0, err
	}

	return infos, total, nil
}
