package implementation

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	api "github.com/Planck1858/golang101/examples/grpc-server/gen/go"
)

type Implementation struct {
	api.UnimplementedGrpcServiceServer

	log   *slog.Logger
	store map[string]Todo
}

func New(log *slog.Logger) *Implementation {
	return &Implementation{
		log:   log,
		store: make(map[string]Todo),
	}
}

func (i *Implementation) RegisterGRPC(server grpc.ServiceRegistrar) {
	api.RegisterGrpcServiceServer(server, i)
}

func (i *Implementation) RegisterHTTP(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	if err := api.RegisterGrpcServiceHandler(ctx, mux, conn); err != nil {
		return fmt.Errorf("register HTTP handler: %w", err)
	}

	return nil
}

func (i *Implementation) GetName() string {
	return "ImplementationService"
}

func (i *Implementation) GetSpec() string {
	return api.SwaggerSpec
}
