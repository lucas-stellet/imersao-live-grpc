package grpc

import "net"

func StartGRPCServer() {
	lis, err := net.Listen("tpc", "localhost:500")
}
