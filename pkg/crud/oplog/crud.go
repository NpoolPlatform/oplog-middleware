package oplog

import (
	"fmt"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/oplog-middleware/pkg/db/ent"
	entoplog "github.com/NpoolPlatform/oplog-middleware/pkg/db/ent/oplog"
	"github.com/google/uuid"
)

type Req struct {
	EntID            *uuid.UUID
	AppID            *uuid.UUID
	UserID           *uuid.UUID
	Path             *string
	Method           *basetypes.HTTPMethod
	Arguments        *string
	CurValue         *string
	NewValue         *string
	HumanReadable    *string
	Result           *basetypes.Result
	FailReason       *string
	ElapsedMillisecs *uint32
}

func CreateSet(c *ent.OpLogCreate, req *Req) *ent.OpLogCreate {
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.Path != nil {
		c.SetPath(*req.Path)
	}
	if req.Method != nil {
		c.SetMethod(req.Method.String())
	}
	if req.Arguments != nil {
		c.SetArguments(*req.Arguments)
	}
	if req.CurValue != nil {
		c.SetCurValue(*req.CurValue)
	}
	if req.HumanReadable != nil {
		c.SetHumanReadable(*req.HumanReadable)
	}
	if req.Result != nil {
		c.SetResult(req.Result.String())
	}
	if req.FailReason != nil {
		c.SetFailReason(*req.FailReason)
	}
	if req.ElapsedMillisecs != nil {
		c.SetElapsedMillisecs(*req.ElapsedMillisecs)
	}
	return c
}

func UpdateSet(u *ent.OpLogUpdateOne, req *Req) *ent.OpLogUpdateOne {
	if req.HumanReadable != nil {
		u.SetHumanReadable(*req.HumanReadable)
	}
	if req.NewValue != nil {
		u.SetNewValue(*req.NewValue)
	}
	if req.Result != nil {
		u.SetResult(req.Result.String())
	}
	if req.FailReason != nil {
		u.SetFailReason(*req.FailReason)
	}
	if req.ElapsedMillisecs != nil {
		u.SetElapsedMillisecs(*req.ElapsedMillisecs)
	}
	return u
}

type Conds struct {
	EntID  *cruder.Cond
	AppID  *cruder.Cond
	UserID *cruder.Cond
	Result *cruder.Cond
}

func SetQueryConds(q *ent.OpLogQuery, conds *Conds) (*ent.OpLogQuery, error) {
	if conds.EntID != nil {
		switch conds.EntID.Op {
		case cruder.EQ:
			id, ok := conds.EntID.Val.(uuid.UUID)
			if !ok {
				return nil, fmt.Errorf("invalid auto id")
			}
			q.Where(entoplog.EntID(id))
		default:
			return nil, fmt.Errorf("invalid oplog field")
		}
	}
	if conds.AppID != nil {
		switch conds.AppID.Op {
		case cruder.EQ:
			id, ok := conds.AppID.Val.(uuid.UUID)
			if !ok {
				return nil, fmt.Errorf("invalid id")
			}
			q.Where(entoplog.AppID(id))
		default:
			return nil, fmt.Errorf("invalid oplog field")
		}
	}
	if conds.UserID != nil {
		switch conds.UserID.Op {
		case cruder.EQ:
			id, ok := conds.UserID.Val.(uuid.UUID)
			if !ok {
				return nil, fmt.Errorf("invalid id")
			}
			q.Where(entoplog.UserID(id))
		default:
			return nil, fmt.Errorf("invalid oplog field")
		}
	}
	if conds.Result != nil {
		switch conds.Result.Op {
		case cruder.EQ:
			result, ok := conds.Result.Val.(basetypes.Result)
			if !ok {
				return nil, fmt.Errorf("invalid oplog col")
			}
			q.Where(entoplog.Result(result.String()))
		default:
			return nil, fmt.Errorf("invalid oplog field")
		}
	}
	q.Where(entoplog.DeletedAt(0))
	return q, nil
}
