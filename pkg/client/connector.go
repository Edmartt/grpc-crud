package client

import (
	"os"

	"github.com/edmartt/grpc-test/internal/utils"
	"google.golang.org/grpc"
)

func grpcConnector() (*grpc.ClientConn, error) {
	zlog := utils.NewZeroLoggerAdapter()
	serverAddress := os.Getenv("SERVER_ADDRESS")

	zlog.Info("starting grpc connection")

	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())

	if err != nil {
		zlog.Error(err.Error())
		return nil, err
	}

	return conn, nil
}
