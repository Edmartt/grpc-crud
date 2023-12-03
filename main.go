package main

import (
	"github.com/edmartt/grpc-test/internal/database"
	"github.com/edmartt/grpc-test/internal/server"
	"github.com/edmartt/grpc-test/pkg/client/http"
	"github.com/joho/godotenv"
)

func main() {

	go http.Start("8080")
	godotenv.Load()
	database.InitMigrations()
	server.StartServer()
}
