package api

import (
	"context"

	"github.com/NpoolPlatform/innovation-minning/message/npool"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedInnovationMinningServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterInnovationMinningServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterInnovationMinningHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
