package http

import (
	"github.com/gin-gonic/gin"

	_ "github.com/edmartt/grpc-test/docs"
)

func setRouter() *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("/api/v1")
	setPersonRoutes(apiGroup)

	return router
}

func Start(port string) {
	router := setRouter()
	router.Run(":" + port)
}
