package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func (s *Server) runGrpcServer(imps []Implementation) (func(), error) {
	// grpc tcp listener
	grpcPortStr := fmt.Sprintf(":%d", s.grpcPort)
	lis, err := net.Listen("tcp", grpcPortStr)
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}

	// initialize grpc server
	grpcServer := s.initGrpcServer(imps)

	// run grpc server
	go func() {
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatalf("Error on start listen gRPC server: %s", err)
		}
	}()

	log.Printf("Serving gRPC on http://0.0.0.0%v", grpcPortStr)

	return grpcServer.GracefulStop, nil
}

func (s *Server) initGrpcServer(imps []Implementation) *grpc.Server {
	// server options
	opts := make([]grpc.ServerOption, 0)
	opts = append(opts, s.grpcServerOptions...)

	// interceptors (aka middlewares)
	ins := make([]grpc.UnaryServerInterceptor, 0)
	ins = append(ins, s.grpcMiddlewares...)
	opts = append(opts, grpc.ChainUnaryInterceptor(ins...))

	// gRPC server object
	serv := grpc.NewServer(opts...)

	// register all implementations
	for _, imp := range imps {
		imp.RegisterGRPC(serv)
	}

	return serv
}
