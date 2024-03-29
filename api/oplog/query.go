package oplog

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/oplog/mw/v1/oplog"
	oplog1 "github.com/NpoolPlatform/oplog-middleware/pkg/mw/oplog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetOpLogs(ctx context.Context, in *npool.GetOpLogsRequest) (*npool.GetOpLogsResponse, error) {
	handler, err := oplog1.NewHandler(
		ctx,
		oplog1.WithConds(in.GetConds()),
		oplog1.WithOffset(in.GetOffset()),
		oplog1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOpLogs",
			"In", in,
			"Error", err,
		)
		return &npool.GetOpLogsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetOpLogs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOpLogs",
			"In", in,
			"Error", err,
		)
		return &npool.GetOpLogsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOpLogsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
