package main

import (
	"os"

	"github.com/edmartt/grpc-test/internal/database"
	"github.com/edmartt/grpc-test/internal/server"
	"github.com/edmartt/grpc-test/pkg/client/http"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("HTTP_PORT")
	go http.Start(port)

	database.InitMigrations()
	server.StartServer()
}
