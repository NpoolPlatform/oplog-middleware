package oplog

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/oplog-middleware/pkg/testinit"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/oplog/mw/v1/oplog"
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
	AppID:       uuid.NewString(),
	UserID:      &userID,
	Path:        uuid.NewString(),
	Method:      basetypes.HTTPMethod_GET,
	Arguments:   "{}",
	CurValue:    `{"A":"a", "B": 18}`,
	ReqHeaders:  "{}",
	RespHeaders: "{}",
}

var req = &npool.OpLogReq{
	AppID:     &ret.AppID,
	UserID:    ret.UserID,
	Path:      &ret.Path,
	Method:    &ret.Method,
	Arguments: &ret.Arguments,
	CurValue:  &ret.CurValue,
}

func create(t *testing.T) {
	info, err := CreateOpLog(context.Background(), req)
	if assert.Nil(t, err) && assert.NotNil(t, info) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.EntID = info.EntID
		ret.MethodStr = info.MethodStr
		ret.ResultStr = info.ResultStr
		assert.Equal(t, ret.String(), info.String())
	}
}

func update(t *testing.T) {
	ret.Result = basetypes.Result_Fail
	ret.FailReason = "I don't know"
	ret.HumanReadable = "Somebody want to kill me"

	req.Result = &ret.Result
	req.FailReason = &ret.FailReason
	req.HumanReadable = &ret.HumanReadable
	req.EntID = &ret.EntID

	info, err := UpdateOpLog(context.Background(), req)
	if assert.Nil(t, err) && assert.NotNil(t, info) {
		ret.UpdatedAt = info.UpdatedAt
		ret.ResultStr = info.ResultStr
		assert.Equal(t, ret.String(), info.String())
	}
}

func getConds(t *testing.T) {
	infos, total, err := GetOpLogs(
		context.Background(),
		&npool.Conds{
			EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
			AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			UserID: &basetypes.StringVal{Op: cruder.EQ, Value: *ret.UserID},
			Result: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.Result)},
		},
		0,
		2,
	)
	if assert.Nil(t, err) {
		assert.Equal(t, uint32(1), total)
		if assert.Equal(t, 1, len(infos)) {
			assert.Equal(t, infos[0].String(), ret.String())
		}
	}
}

func TestOpLog(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)
	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("create", create)
	t.Run("update", update)
	t.Run("getConds", getConds)
}
