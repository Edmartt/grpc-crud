package client

import (
	"log"
	"os"

	"google.golang.org/grpc"
)

func grpcConnector() *grpc.ClientConn {
	serverAddress := os.Getenv("SERVER_ADDRESS")
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())

	if err != nil {
		log.Println(err.Error())
	}

	return conn
}
