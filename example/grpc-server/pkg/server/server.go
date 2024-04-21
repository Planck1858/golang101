package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

const (
	versionTimeFormat = "02-01-2006 03:04:05"
)

type (
	Server struct {
		grpcPort         int
		httpPort         int
		maxMsgSize       int
		readTimeout      int
		writeTimeout     int
		enablePprof      bool
		versionCreatedAt time.Time
		versionTag       string

		ctx    context.Context
		cancel context.CancelFunc

		grpcMiddlewares               []grpc.UnaryServerInterceptor
		grpcServerOptions             []grpc.ServerOption
		gatewayMiddlewares            []Handler
		gatewayToGrpcCallOptions      []grpc.CallOption
		gatewayForwardResponseOptions []ForwardResponseFunc
		gatewayIncomingHeaderMatchers []IncomingHeaderFunc
	}

	Implementation interface {
		RegisterGRPC(server grpc.ServiceRegistrar)
		RegisterHTTP(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
		GetSpec() string
		GetName() string
	}

	Handler             func(h http.Handler) http.Handler
	ForwardResponseFunc func(context.Context, http.ResponseWriter, proto.Message) error
	IncomingHeaderFunc  func(key string) (string, bool)
)

func NewServer(otp ...option) *Server {
	s := &Server{}

	for _, o := range otp {
		o(s)
	}

	return s
}

func (s *Server) Run(imps []Implementation) error {
	if s.grpcPort == 0 {
		return fmt.Errorf("grpc port must be set")
	}

	s.ctx, s.cancel = context.WithCancel(context.Background())

	// run grpc server
	stopGrpcServ, err := s.runGrpcServer(imps)
	if err != nil {
		return fmt.Errorf("run gRPC server: %w", err)
	}

	// run http (gRPC GW) server
	var stopGrpcGWServ func(ctx context.Context) error
	if s.httpPort > 0 {
		stopGrpcGWServ, err = s.runGrpcHTTPServer(imps)
		if err != nil {
			return fmt.Errorf("run gRPC GW server: %w", err)
		}
	}

	<-s.ctx.Done()

	// stop grpc server
	stopGrpcServ()
	if s.httpPort > 0 {
		_ = stopGrpcGWServ(context.Background())
	}

	log.Printf("gRPC and gRPC GW servers stopped")

	return nil
}

func (s *Server) Stop() {
	s.cancel()
}
