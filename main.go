package main

import (
	"github.com/edmartt/grpc-test/internal/database"
	"github.com/edmartt/grpc-test/internal/server"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	database.InitMigrations()
	server.StartServer()
}
