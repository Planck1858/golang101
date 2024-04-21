package main

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	"github.com/Planck1858/golang101/examples/grpc-server/internal/implementation"
	"github.com/Planck1858/golang101/examples/grpc-server/pkg/server"
	servergrpc "github.com/Planck1858/golang101/examples/grpc-server/pkg/server/grpc"
	serverhttp "github.com/Planck1858/golang101/examples/grpc-server/pkg/server/http"
	serverinterceptors "github.com/Planck1858/golang101/examples/grpc-server/pkg/server/interceptors"
	servermiddlewares "github.com/Planck1858/golang101/examples/grpc-server/pkg/server/middlewares"
)

const (
	grpcPort = 9000
	httpPort = 80

	maxMsgSize = 1024 * 1024
)

func main() {
	// implementations
	imp := implementation.New(slog.Default())

	// grpc server
	s := server.NewServer(
		server.GRPCPort(grpcPort),
		server.HTTPPort(httpPort),
		server.ReadWriteTimeoutInSec(15, 15),
		server.MaxMsgSize(1024*1024),
		server.EnablePprof(true),
		server.GRPCServOptions([]grpc.ServerOption{
			servergrpc.MaxReceiveMsgSize(maxMsgSize),
			servergrpc.MaxSendMsgSize(maxMsgSize),
		}),
		server.GRPCCallOptions([]grpc.CallOption{
			serverhttp.MaxSendMsgSize(maxMsgSize),
			serverhttp.MaxReceiveMsgSize(maxMsgSize),
		}),
		server.GatewayMiddlewares([]server.Handler{
			servermiddlewares.NewRequestIDMiddleware(),
		}),
		server.GRPCInterceptors([]grpc.UnaryServerInterceptor{
			serverinterceptors.RequestIDInterceptor(),
		}),
		server.IncomingHeaderMatchers([]server.IncomingHeaderFunc{
			GrpcGwCustomHeaderMatcher,
		}),
		// server.ForwardResponseOptions([]server.ForwardResponseFunc{
		// 	GrpcGwResponseModifier,
		// }),
		server.Version("v1.0.0", time.Now()),
	)

	// run grpc server
	err := s.Run([]server.Implementation{
		imp,
	})
	if err != nil {
		panic(err)
	}
}

/* Custom header matcher */
const (
	allowKey = "allow"
	denyKey  = "deny"
)

func GrpcGwCustomHeaderMatcher(key string) (string, bool) {
	switch key {
	case allowKey:
		return key, true
	case denyKey:
		return "", false
	}

	return key, true
}

/* Forward response function */
// WithForwardResponseOption returns a ServeMuxOption representing the forwardResponseOption.
//
// forwardResponseOption is an option that will be called on the relevant context.Context,
// http.ResponseWriter, and proto.Message before every forwarded response.

// GrpcGwResponseModifier - метод работает как middleware на выходе всех запросов grpc gw -> grpc.
// Позволяет установить смайлик в каждый ответ grpc-gw -> grpc
func GrpcGwResponseModifier(ctx context.Context, rw http.ResponseWriter, p proto.Message) error {
	rw.Write([]byte("\n:)"))
	return nil
}
