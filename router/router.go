package router

import (
	"util-go/docs"
	"util-go/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/ping", service.Ping)
	r.GET("/getAuthorization", service.GetAuthorization)
	r.GET("/parseAuthorization", service.ParseAuthorization)

	return r
}
