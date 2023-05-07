package oplog

import (
	"fmt"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/oplog-middleware/pkg/db/ent"
	entoplog "github.com/NpoolPlatform/oplog-middleware/pkg/db/ent/oplog"
	"github.com/google/uuid"
)

type Req struct {
	ID        *uuid.UUID
	SampleCol *string
}

func CreateSet(c *ent.OpLogCreate, req *Req) *ent.OpLogCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.SampleCol != nil {
		c.SetSampleCol(*req.SampleCol)
	}
	return c
}

func UpdateSet(u *ent.OpLogUpdateOne, req *Req) *ent.OpLogUpdateOne {
	if req.SampleCol != nil {
		u.SetSampleCol(*req.SampleCol)
	}
	return u
}

type Conds struct {
	AutoID    *cruder.Cond
	ID        *cruder.Cond
	SampleCol *cruder.Cond
}

func SetQueryConds(q *ent.OpLogQuery, conds *Conds) (*ent.OpLogQuery, error) {
	if conds.AutoID != nil {
		switch conds.AutoID.Op {
		case cruder.EQ:
			id, ok := conds.AutoID.Val.(uint32)
			if !ok {
				return nil, fmt.Errorf("invalid auto id")
			}
			q.Where(entoplog.AutoID(id))
		default:
			return nil, fmt.Errorf("invalid oplog field")
		}
	}
	if conds.ID != nil {
		switch conds.ID.Op {
		case cruder.EQ:
			id, ok := conds.ID.Val.(uuid.UUID)
			if !ok {
				return nil, fmt.Errorf("invalid id")
			}
			q.Where(entoplog.ID(id))
		default:
			return nil, fmt.Errorf("invalid oplog field")
		}
	}
	if conds.SampleCol != nil {
		switch conds.SampleCol.Op {
		case cruder.LIKE:
			sampleCol, ok := conds.ID.Val.(string)
			if !ok {
				return nil, fmt.Errorf("invalid oplog col")
			}
			q.Where(entoplog.SampleCol(sampleCol))
		default:
			return nil, fmt.Errorf("invalid oplog field")
		}
	}
	q.Where(entoplog.DeletedAt(0))
	return q, nil
}
