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
	if req == nil {
		logger.Sugar().Errorw(
			"CreateOpLog",
			"In", in,
		)
		return &npool.CreateOpLogResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := oplog1.NewHandler(
		ctx,
		oplog1.WithAppID(req.GetAppID(), true),
		oplog1.WithUserID(req.UserID, false),
		oplog1.WithPath(req.Path, true),
		oplog1.WithMethod(req.Method, false),
		oplog1.WithArguments(req.Arguments, false),
		oplog1.WithCurValue(req.CurValue, false),
		oplog1.WithHumanReadable(req.HumanReadable, false),
		oplog1.WithReqHeaders(req.ReqHeaders, false),
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
