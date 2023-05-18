//nolint:dupl
package oplog

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/oplog/mw/v1/oplog"

	servicename "github.com/NpoolPlatform/oplog-middleware/pkg/servicename"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(servicename.ServiceDomain, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get oplog connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func CreateOpLog(ctx context.Context, in *npool.OpLogReq) (*npool.OpLog, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateOpLog(ctx, &npool.CreateOpLogRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create oplog: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create oplog: %v", err)
	}
	return info.(*npool.OpLog), nil
}

func UpdateOpLog(ctx context.Context, in *npool.OpLogReq) (*npool.OpLog, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateOpLog(ctx, &npool.UpdateOpLogRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update oplog: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update oplog: %v", err)
	}
	return info.(*npool.OpLog), nil
}

func GetOpLogs(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.OpLog, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetOpLogs(ctx, &npool.GetOpLogsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get oplogs: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get oplogs: %v", err)
	}
	return infos.([]*npool.OpLog), total, nil
}
