package api

import (
	"context"

	oplog "github.com/NpoolPlatform/message/npool/oplog/mw/v1"
	oplog1 "github.com/NpoolPlatform/oplog-middleware/api/oplog"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	oplog.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	oplog.RegisterMiddlewareServer(server, &Server{})
	oplog1.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := oplog.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := oplog1.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
