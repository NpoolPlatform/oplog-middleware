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
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateOpLog",
			"In", in,
		)
		return &npool.UpdateOpLogResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := oplog1.NewHandler(
		ctx,
		oplog1.WithID(req.ID, true),
		oplog1.WithNewValue(req.NewValue, false),
		oplog1.WithHumanReadable(req.HumanReadable, false),
		oplog1.WithResult(req.Result, false),
		oplog1.WithFailReason(req.FailReason, false),
		oplog1.WithStatusCode(req.StatusCode, false),
		oplog1.WithReqHeaders(req.ReqHeaders, false),
		oplog1.WithRespHeaders(req.RespHeaders, false),
		oplog1.WithElapsedMillisecs(req.GetElapsedMillisecs(), false),
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
