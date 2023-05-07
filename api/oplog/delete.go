package oplog

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/oplog/mw/v1/oplog"
	oplog1 "github.com/NpoolPlatform/oplog-middleware/pkg/mw/oplog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteOpLog(ctx context.Context, in *npool.DeleteOpLogRequest) (*npool.DeleteOpLogResponse, error) {
	req := in.GetInfo()
	handler, err := oplog1.NewHandler(
		ctx,
		oplog1.WithID(ctx, req.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteOpLog",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteOpLogResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteOpLog(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteOpLog",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteOpLogResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteOpLogResponse{
		Info: info,
	}, nil
}
