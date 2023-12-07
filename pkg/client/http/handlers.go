package http

import (
	"net/http"

	"github.com/edmartt/grpc-test/internal/person/models"
	pb "github.com/edmartt/grpc-test/internal/person/protos/bin"
	"github.com/edmartt/grpc-test/pkg/client"
	"github.com/gin-gonic/gin"
)

//	@title			HTTP Client for gRPC Client
//	@version		1.0
//	@description	This client handles data for sending data to gRPC client and after that to gRPC server
//	@termsOfService	http://swagger.io/terms/

// @host		localhost:8080
// @BasePath	/api/v1

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

	response, err := client.ReadPerson(requestPB)

	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	}

	if response.Person.Id == "" {
		context.JSON(http.StatusNotFound, httpResponse{
			Response: "Not found",
		})
		return
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

	response, err := client.CreatePerson(personProtoModel)

	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	}

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
		return
	}

	response, err := client.DeletePerson(requestPB)

	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	if response.Status == "not found" {
		context.JSON(http.StatusNotFound, httpResponse{
			Response: "Not found",
		})
		return
	}

	context.JSON(http.StatusOK, response)
}
