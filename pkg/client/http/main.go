package main

import (
	"log"
	"net/http"

	"github.com/edmartt/grpc-test/internal/person/models"
	pb "github.com/edmartt/grpc-test/internal/person/protos/bin"
	"github.com/edmartt/grpc-test/pkg/client"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type createdResponse struct {
	Response string `json:"response"`
}

func getPerson(context *gin.Context) {
	id := context.Param("id")

	requestPB := &pb.GetPersonRequest{
		Id: id,
	}

	response := client.ReadPerson(requestPB)
	log.Println("response client grpc: ", response)

	context.JSON(http.StatusOK, response)
}

func postPerson(context *gin.Context) {
	personModel := models.Person{}
	personProtoModel := &pb.Person{}

	err := context.BindJSON(&personModel)

	personProtoModel.FirstName = personModel.FirstName
	personProtoModel.LastName = personModel.LastName
	personProtoModel.Email = personModel.Email

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response := client.CreatePerson(personProtoModel)
	created := createdResponse{
		Response: response,
	}

	context.JSON(http.StatusOK, created)
}

func deletePerson(context *gin.Context) {
	id := context.Param("id")
	requestPB := &pb.DeletePersonRequest{
		Id: id,
	}

	response := client.DeletePerson(requestPB)

	context.JSON(http.StatusOK, response)
}

func setPersonRoutes(router *gin.RouterGroup) {
	router.POST("/person", postPerson)
	router.GET("/person/:id", getPerson)
	router.DELETE("/person/:id", deletePerson)
}

func setRouter() *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("/api/v1")
	setPersonRoutes(apiGroup)

	return router
}

func start(port string) {
	router := setRouter()
	router.Run(":" + port)
}

func main() {
	godotenv.Load()
	start("8080")
}
