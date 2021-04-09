package main

import (
	"log"
	"net"

	server "example.com/user_name/sample/services/server/grpc"
	pb "example.com/user_name/sample/services/server/pb/picture"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/reflection"
)

const port = ":50050"

func run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return errors.Wrap(err, "port failed")
	}

	grpcServer := grpc.NewServer()
	var svr server.Server
	pb.RegisterPictureServer(grpcServer, &svr)
	// reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		return errors.Wrap(err, "error server start failed")
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("%v", err)
	}
}
