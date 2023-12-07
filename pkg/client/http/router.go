package http

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setPersonRoutes(router *gin.RouterGroup) {
	router.POST("/person", postPerson)
	router.GET("/person/:id", getPerson)
	router.DELETE("/person/:id", deletePerson)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
