package oplog

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	// "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

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
}

var (
	req = &npool.OpLogReq{
		AppID:     &ret.AppID,
		UserID:    ret.UserID,
		Path:      &ret.Path,
		Method:    &ret.Method,
		Arguments: &ret.Arguments,
	}
)

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithAppID(context.Background(), ret.AppID),
		WithUserID(context.Background(), req.UserID),
		WithPath(context.Background(), req.Path),
		WithMethod(context.Background(), req.Method),
		WithArguments(context.Background(), req.Arguments),
		WithCurValue(context.Background(), req.CurValue),
		WithHumanReadable(context.Background(), req.HumanReadable),
		WithResult(context.Background(), req.Result),
		WithFailReason(context.Background(), req.FailReason),
	)
	if assert.Nil(t, err) {
		info, err := handler.CreateOpLog(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.AutoID = info.AutoID
			assert.Equal(t, info, ret)
		}
	}
}

/*
func update(t *testing.T) {
	depracated := true

	req.Depracated = &depracated
	ret.Depracated = depracated

	info, err := Update(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), ret.String())
	}
}

func getConds(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		if assert.Equal(t, total, 1) {
			assert.Equal(t, infos[0].String(), ret.String())
		}
	}
}
*/

func TestAPI(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	// t.Run("update", update)
	// t.Run("getConds", getConds)
}
