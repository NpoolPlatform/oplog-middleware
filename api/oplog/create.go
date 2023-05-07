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

func (s *Server) CreateOpLog(ctx context.Context, in *npool.CreateOpLogRequest) (*npool.CreateOpLogResponse, error) {
	req := in.GetInfo()
	handler, err := oplog1.NewHandler(
		ctx,
		oplog1.WithID(ctx, req.ID),
		oplog1.WithSampleCol(ctx, req.SampleCol),
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

func (s *Server) CreateOpLogs(ctx context.Context, in *npool.CreateOpLogsRequest) (*npool.CreateOpLogsResponse, error) {
	handler, err := oplog1.NewHandler(
		ctx,
		oplog1.WithReqs(ctx, in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateOpLogs",
			"In", in,
			"Error", err,
		)
		return &npool.CreateOpLogsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, err := handler.CreateOpLogs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateOpLogs",
			"In", in,
			"Error", err,
		)
		return &npool.CreateOpLogsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateOpLogsResponse{
		Infos: infos,
	}, nil
}
