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

func (s *Server) ExistOpLog(ctx context.Context, in *npool.ExistOpLogRequest) (*npool.ExistOpLogResponse, error) {
	handler, err := oplog1.NewHandler(
		ctx,
		oplog1.WithID(ctx, &in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistOpLog",
			"In", in,
			"Error", err,
		)
		return &npool.ExistOpLogResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := handler.ExistOpLog(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistOpLog",
			"In", in,
			"Error", err,
		)
		return &npool.ExistOpLogResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistOpLogResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistOpLogConds(ctx context.Context, in *npool.ExistOpLogCondsRequest) (*npool.ExistOpLogCondsResponse, error) {
	handler, err := oplog1.NewHandler(
		ctx,
		oplog1.WithConds(ctx, in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistOpLogConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistOpLogCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := handler.ExistOpLogConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistOpLogConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistOpLogCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistOpLogCondsResponse{
		Info: exist,
	}, nil
}
