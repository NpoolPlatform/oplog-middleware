package oplog

import (
	"github.com/NpoolPlatform/message/npool/oplog/mw/v1/oplog"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	oplog.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	oplog.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
