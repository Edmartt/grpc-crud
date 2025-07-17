package server

import (
	"net"

	"github.com/edmartt/grpc-test/internal/person"
	pb "github.com/edmartt/grpc-test/internal/person/protos/bin"
	"github.com/edmartt/grpc-test/internal/utils"
	"google.golang.org/grpc"
)

func StartServer(port string) {

	listener, err := net.Listen("tcp", ":"+port)

	zlog := utils.NewZeroLoggerAdapter()

	if err != nil {
		zlog.Fatal(err.Error())
	}

	serve := grpc.NewServer()

	zlog.Info("SERVER RUNNING on: " + port)

	pb.RegisterPersonServiceServer(serve, &person.Service{})

	if err = serve.Serve(listener); err != nil {
		zlog.Error("server not started " + err.Error())
	}
}
