package server

import (
	"time"

	"google.golang.org/grpc"
)

type option func(conf *Server)

func GRPCPort(port int) option {
	return func(s *Server) {
		s.grpcPort = port
	}
}

func HTTPPort(port int) option {
	return func(s *Server) {
		s.httpPort = port
	}
}

func ReadWriteTimeoutInSec(read, write int) option {
	return func(s *Server) {
		s.readTimeout = read
		s.writeTimeout = write
	}
}

func MaxMsgSize(size int) option {
	return func(s *Server) {
		s.maxMsgSize = size
	}
}

func EnablePprof(en bool) option {
	return func(s *Server) {
		s.enablePprof = en
	}
}

func GRPCServOptions(grpcServOptions []grpc.ServerOption) option {
	return func(s *Server) {
		s.grpcServerOptions = grpcServOptions
	}
}

func GRPCCallOptions(grpcCallOptions []grpc.CallOption) option {
	return func(s *Server) {
		s.gatewayToGrpcCallOptions = grpcCallOptions
	}
}

func GatewayMiddlewares(gatewayMiddlewares []Handler) option {
	return func(s *Server) {
		s.gatewayMiddlewares = gatewayMiddlewares
	}
}

func GRPCInterceptors(grpcInterceptors []grpc.UnaryServerInterceptor) option {
	return func(s *Server) {
		s.grpcMiddlewares = grpcInterceptors
	}
}

func ForwardResponseOptions(forwardResponseOptions []ForwardResponseFunc) option {
	return func(s *Server) {
		s.gatewayForwardResponseOptions = forwardResponseOptions
	}
}

func IncomingHeaderMatchers(incomingHeaderMatchers []IncomingHeaderFunc) option {
	return func(s *Server) {
		s.gatewayIncomingHeaderMatchers = incomingHeaderMatchers
	}
}

func Version(version string, date time.Time) option {
	return func(s *Server) {
		s.versionTag = version
		s.versionCreatedAt = date
	}
}
