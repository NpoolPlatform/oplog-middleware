package oplog

import (
	"context"
	"encoding/json"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/oplog/mw/v1/oplog"
	constant "github.com/NpoolPlatform/oplog-middleware/pkg/const"
	oplogcrud "github.com/NpoolPlatform/oplog-middleware/pkg/crud/oplog"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID               *uint32
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
	StatusCode       *int32
	ReqHeaders       *string
	RespHeaders      *string
	ElapsedMillisecs *uint32
	Conds            *oplogcrud.Conds
	Offset           int32
	Limit            int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = u
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = &_id
		return nil
	}
}

func WithAppID(id string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_id, err := uuid.Parse(id)
		if err != nil {
			return err
		}
		h.AppID = &_id
		return nil
	}
}

func WithUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid userid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.UserID = &_id
		return nil
	}
}

func WithPath(path *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Path = path
		return nil
	}
}

func WithMethod(method *basetypes.HTTPMethod, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if method == nil {
			if must {
				return fmt.Errorf("invalid method")
			}
			return nil
		}
		switch *method {
		case basetypes.HTTPMethod_GET:
		case basetypes.HTTPMethod_HEAD:
		case basetypes.HTTPMethod_POST:
		case basetypes.HTTPMethod_PUT:
		case basetypes.HTTPMethod_DELETE:
		case basetypes.HTTPMethod_CONNECT:
		case basetypes.HTTPMethod_OPTIONS:
		case basetypes.HTTPMethod_TRACE:
		case basetypes.HTTPMethod_PATCH:
		default:
			return fmt.Errorf("invalid method")
		}
		h.Method = method
		return nil
	}
}

func WithArguments(args *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if args == nil {
			if must {
				return fmt.Errorf("invalid arguments")
			}
			return nil
		}
		var _args map[string]interface{}
		if err := json.Unmarshal([]byte(*args), &_args); err != nil {
			return err
		}
		h.Arguments = args
		return nil
	}
}

func WithCurValue(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid curvalue")
			}
			return nil
		}
		var _args map[string]interface{}
		if err := json.Unmarshal([]byte(*value), &_args); err != nil {
			return err
		}
		h.CurValue = value
		return nil
	}
}

func WithNewValue(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid newvalue")
			}
			return nil
		}
		var _args map[string]interface{}
		if err := json.Unmarshal([]byte(*value), &_args); err != nil {
			return err
		}
		h.NewValue = value
		return nil
	}
}

func WithHumanReadable(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.HumanReadable = value
		return nil
	}
}

func WithResult(result *basetypes.Result, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if result == nil {
			if must {
				return fmt.Errorf("invalid result")
			}
			return nil
		}
		switch *result {
		case basetypes.Result_Success:
		case basetypes.Result_Fail:
		default:
			return fmt.Errorf("invalid result")
		}
		h.Result = result
		return nil
	}
}

func WithFailReason(reason *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.FailReason = reason
		return nil
	}
}

func WithStatusCode(statusCode *int32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.StatusCode = statusCode
		return nil
	}
}

func WithReqHeaders(headers *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.ReqHeaders = headers
		return nil
	}
}

func WithRespHeaders(headers *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.RespHeaders = headers
		return nil
	}
}

func WithElapsedMillisecs(millisecs uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.ElapsedMillisecs = &millisecs
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &oplogcrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			h.Conds.ID = &cruder.Cond{Op: conds.GetID().GetOp(), Val: conds.GetID().GetValue()}
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{
				Op:  conds.GetEntID().GetOp(),
				Val: id,
			}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{
				Op:  conds.GetAppID().GetOp(),
				Val: id,
			}
		}
		if conds.UserID != nil {
			id, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{
				Op:  conds.GetUserID().GetOp(),
				Val: id,
			}
		}
		if conds.Result != nil {
			h.Conds.Result = &cruder.Cond{
				Op:  conds.GetResult().GetOp(),
				Val: basetypes.Result(conds.GetResult().GetValue()),
			}
		}
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
