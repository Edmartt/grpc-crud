package main

import (
	"net/http"

	"github.com/edmartt/grpc-test/internal/person/models"
	pb "github.com/edmartt/grpc-test/internal/person/protos/bin"
	"github.com/edmartt/grpc-test/pkg/client"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/edmartt/grpc-test/docs"
)

//	@title			HTTP Client for gRPC Client
//	@version		1.0
//	@description	This client handles data for sending data to gRPC client and after that to gRPC server
//	@termsOfService	http://swagger.io/terms/
//
// @host		localhost:8080
// @BasePath	/api/v1
type httpResponse struct {
	Response string `json:"response"`
}

// getPerson godoc
// @Summary      Get persons from DB
// @Description  Through a get request the id is sent to gRPC client
// @Tags         Persons
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Person ID"
// @Success      200  {object}  models.Person
// @Failure	 400 {object}  httpResponse
// @Failure	 404 {object} httpResponse
// @Router       /person/{id} [get]
func getPerson(context *gin.Context) {
	id := context.Param("id")

	badReq := httpResponse{
		Response: "ID empty",
	}

	if id == "" {
		context.JSON(http.StatusBadRequest, badReq)
	}

	requestPB := &pb.GetPersonRequest{
		Id: id,
	}

	response := client.ReadPerson(requestPB)

	if response.Person.Id == "" {
		context.JSON(http.StatusNotFound, httpResponse{
			Response: "Not found",
		})
	}

	context.JSON(http.StatusOK, response)
}

// postPerson godoc
// @Summary      Creates new person
// @Description  This endpoint is for creating persons
// @Tags         Persons
// @Accept       json
// @Produce      json
// @Param        person body  models.Person true "Creates person"
// @Success      200  {object}  httpResponse
// @Failure	 400 {object}  httpResponse
// @Router       /person [post]
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
	created := httpResponse{
		Response: response,
	}

	context.JSON(http.StatusOK, created)
}

// deletePerson godoc
// @Summary      Deletes person by ID
// @Description  This endpoint is for deleting person by ID
// @Tags         Persons
// @Accept       json
// @Produce      json
// @Param        id  path string true "uuid formatted ID"
// @Success      200  {object}  httpResponse
// @Failure	 400 {object}  httpResponse
// @Failure	 404 {object} httpResponse
// @Router       /person/{id} [delete]
func deletePerson(context *gin.Context) {
	id := context.Param("id")
	requestPB := &pb.DeletePersonRequest{
		Id: id,
	}

	if id == "" {
		context.JSON(http.StatusBadRequest, httpResponse{
			Response: "ID is missing",
		})
	}

	response := client.DeletePerson(requestPB)

	if response.Id == "" {
		context.JSON(http.StatusNotFound, httpResponse{
			Response: "Not found",
		})
	}

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
