package http

import "google.golang.org/grpc"

func MaxSendMsgSize(size int) grpc.CallOption {
	return grpc.MaxCallSendMsgSize(size)
}

func MaxReceiveMsgSize(size int) grpc.CallOption {
	return grpc.MaxCallRecvMsgSize(size)
}
