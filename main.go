package main

import (
	"log"
	"os"

	"github.com/edmartt/grpc-test/internal/database"
	"github.com/edmartt/grpc-test/internal/server"
	"github.com/edmartt/grpc-test/pkg/client/http"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("HTTP_PORT")
	grpcPort := os.Getenv("PORT")
	go http.Start(port)

	database.InitMigrations()
	server.StartServer(grpcPort)
}
