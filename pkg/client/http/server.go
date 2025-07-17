package http

import (
	"github.com/gin-gonic/gin"

	_ "github.com/edmartt/grpc-test/docs"
	"github.com/edmartt/grpc-test/internal/utils"
)

func setRouter() *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("/api/v1")
	setPersonRoutes(apiGroup)

	return router
}

func Start(port string) {
	router := setRouter()
	zlog := utils.NewZeroLoggerAdapter()
	err := router.Run(":" + port)

	if err != nil {
		zlog.Error(err.Error())
	}
}
