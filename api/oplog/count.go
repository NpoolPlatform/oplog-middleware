//nolint:dupl
package oplog

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/oplog/mw/v1/oplog"
	oplog1 "github.com/NpoolPlatform/oplog-middleware/pkg/mw/oplog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CountOpLogs(ctx context.Context, in *npool.CountOpLogsRequest) (*npool.CountOpLogsResponse, error) {
	handler, err := oplog1.NewHandler(
		ctx,
		oplog1.WithConds(ctx, in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CountOpLogs",
			"In", in,
			"Error", err,
		)
		return &npool.CountOpLogsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	count, err := handler.CountOpLogs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CountOpLogs",
			"In", in,
			"Error", err,
		)
		return &npool.CountOpLogsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountOpLogsResponse{
		Info: count,
	}, nil
}
