package main

import (
	"os"

	"github.com/edmartt/grpc-test/internal/database"
	"github.com/edmartt/grpc-test/internal/server"
	"github.com/edmartt/grpc-test/internal/utils"
	"github.com/edmartt/grpc-test/pkg/client/http"
	"github.com/joho/godotenv"
)

func main() {
	zLog := utils.NewZeroLoggerAdapter()

	err := godotenv.Load()

	if err != nil {
		zLog.Fatal(err.Error())
	}

	port := os.Getenv("HTTP_PORT")
	grpcPort := os.Getenv("PORT")
	go http.Start(port)

	database.InitMigrations()
	server.StartServer(grpcPort)
}
