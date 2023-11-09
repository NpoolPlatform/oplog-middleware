package oplog

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/oplog/mw/v1/oplog"
	oplogcrud "github.com/NpoolPlatform/oplog-middleware/pkg/crud/oplog"
	"github.com/NpoolPlatform/oplog-middleware/pkg/db"
	"github.com/NpoolPlatform/oplog-middleware/pkg/db/ent"

	"strings"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) checkPath() error {
	if h.Path == nil {
		return nil
	}
	const leastPathUnits = 3
	path := strings.TrimPrefix(*h.Path, "/")
	strs := strings.Split(path, "/")
	if len(strs) < leastPathUnits {
		return fmt.Errorf("invalid path")
	}
	if !strings.HasPrefix(strs[2], "v") {
		return fmt.Errorf("invalid path")
	}
	return nil
}

func (h *Handler) CreateOpLog(ctx context.Context) (*npool.OpLog, error) {
	handler := &createHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.checkPath(); err != nil {
			return err
		}
		info, err := oplogcrud.CreateSet(
			cli.OpLog.Create(),
			&oplogcrud.Req{
				AppID:      h.AppID,
				UserID:     h.UserID,
				Path:       h.Path,
				Method:     h.Method,
				Arguments:  h.Arguments,
				CurValue:   h.CurValue,
				ReqHeaders: h.ReqHeaders,
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
