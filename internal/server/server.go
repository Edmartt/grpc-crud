package server

import (
	"fmt"
	"log"
	"net"

	"github.com/edmartt/grpc-test/internal/person"
	pb "github.com/edmartt/grpc-test/internal/person/protos/bin"
	"google.golang.org/grpc"
)

func StartServer(port string) {

	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		panic(err)
	}

	serve := grpc.NewServer()
	fmt.Println("SERVER RUNNING on: ", port)

	pb.RegisterPersonServiceServer(serve, &person.Service{})

	if err = serve.Serve(listener); err != nil {
		log.Println("Server Not Started " + err.Error())
	}
}
