package oplog

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/oplog/mw/v1/oplog"
	testinit "github.com/NpoolPlatform/oplog-middleware/pkg/testinit"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var userID = uuid.NewString()
var ret = &npool.OpLog{
	AppID:     uuid.NewString(),
	UserID:    &userID,
	Path:      uuid.NewString(),
	Method:    basetypes.HTTPMethod_GET,
	Arguments: "{}",
	CurValue:  `{"A":"a", "B": 18}`,
}

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithAppID(context.Background(), ret.AppID),
		WithUserID(context.Background(), ret.UserID),
		WithPath(context.Background(), &ret.Path),
		WithMethod(context.Background(), &ret.Method),
		WithArguments(context.Background(), &ret.Arguments),
		WithCurValue(context.Background(), &ret.CurValue),
		WithHumanReadable(context.Background(), &ret.HumanReadable),
	)
	if assert.Nil(t, err) {
		info, err := handler.CreateOpLog(context.Background())
		if assert.Nil(t, err) && assert.NotNil(t, info) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.EntID = info.EntID
			ret.MethodStr = info.MethodStr
			ret.ResultStr = info.ResultStr
			assert.Equal(t, ret.String(), info.String())
		}
	}
}

func get(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(context.Background(), &ret.EntID),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetOpLog(context.Background())
		if assert.Nil(t, err) && assert.NotNil(t, info) {
			assert.Equal(t, ret.String(), info.String())
		}
	}
}

func update(t *testing.T) {
	ret.Result = basetypes.Result_Fail
	ret.FailReason = "I don't know"
	ret.HumanReadable = "Somebody want to kill me"

	handler, err := NewHandler(
		context.Background(),
		WithEntID(context.Background(), &ret.EntID),
		WithCurValue(context.Background(), &ret.CurValue),
		WithHumanReadable(context.Background(), &ret.HumanReadable),
		WithResult(context.Background(), &ret.Result),
		WithFailReason(context.Background(), &ret.FailReason),
	)
	if assert.Nil(t, err) {
		info, err := handler.UpdateOpLog(context.Background())
		if assert.Nil(t, err) && assert.NotNil(t, info) {
			ret.UpdatedAt = info.UpdatedAt
			ret.ResultStr = info.ResultStr
			assert.Equal(t, ret.String(), info.String())
		}
	}
}

func getConds(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(context.Background(), &npool.Conds{
			EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
			AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			UserID: &basetypes.StringVal{Op: cruder.EQ, Value: *ret.UserID},
			Result: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.Result)},
		}),
		WithOffset(context.Background(), 0),
		WithLimit(context.Background(), 2),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetOpLogs(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, uint32(1), total)
			if assert.Equal(t, 1, len(infos)) {
				assert.Equal(t, infos[0].String(), ret.String())
			}
		}
	}
}

func TestAPI(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("get", get)
	t.Run("update", update)
	t.Run("getConds", getConds)
}
