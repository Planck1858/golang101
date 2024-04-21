package grpc

import "google.golang.org/grpc"

func MaxSendMsgSize(size int) grpc.ServerOption {
	return grpc.MaxSendMsgSize(size)
}

func MaxReceiveMsgSize(size int) grpc.ServerOption {
	return grpc.MaxRecvMsgSize(size)
}
