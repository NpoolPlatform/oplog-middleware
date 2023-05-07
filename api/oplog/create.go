package oplog

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/oplog/mw/v1/oplog"
	oplog1 "github.com/NpoolPlatform/oplog-middleware/pkg/mw/oplog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateOpLog(ctx context.Context, in *npool.CreateOpLogRequest) (*npool.CreateOpLogResponse, error) {
	req := in.GetInfo()
	handler, err := oplog1.NewHandler(
		ctx,
		oplog1.WithAppID(ctx, req.GetAppID()),
		oplog1.WithUserID(ctx, req.UserID),
		oplog1.WithPath(ctx, req.Path),
		oplog1.WithMethod(ctx, req.Method),
		oplog1.WithArguments(ctx, req.Arguments),
		oplog1.WithCurValue(ctx, req.CurValue),
		oplog1.WithHumanReadable(ctx, req.HumanReadable),
		oplog1.WithResult(ctx, req.Result),
		oplog1.WithFailReason(ctx, req.FailReason),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateOpLog",
			"In", in,
			"Error", err,
		)
		return &npool.CreateOpLogResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateOpLog(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateOpLog",
			"In", in,
			"Error", err,
		)
		return &npool.CreateOpLogResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateOpLogResponse{
		Info: info,
	}, nil
}
