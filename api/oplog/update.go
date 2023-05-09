package oplog

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/oplog/mw/v1/oplog"
	oplog1 "github.com/NpoolPlatform/oplog-middleware/pkg/mw/oplog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateOpLog(ctx context.Context, in *npool.UpdateOpLogRequest) (*npool.UpdateOpLogResponse, error) {
	req := in.GetInfo()
	handler, err := oplog1.NewHandler(
		ctx,
		oplog1.WithEntID(ctx, req.EntID),
		oplog1.WithNewValue(ctx, req.NewValue),
		oplog1.WithHumanReadable(ctx, req.HumanReadable),
		oplog1.WithResult(ctx, req.Result),
		oplog1.WithFailReason(ctx, req.FailReason),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateOpLog",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateOpLogResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateOpLog(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateOpLog",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateOpLogResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateOpLogResponse{
		Info: info,
	}, nil
}
