package route

import (
	"elf-go/example/app/http/handler"
	"elf-go/example/app/http/middelware"
	"github.com/gin-gonic/gin"
)

func InitRoute(router *gin.Engine){
	router.Use(gin.Recovery())
	router.Use(middelware.Cors)

	router.GET("/get-user-info", handler.GetUserInfo)
	router.GET("/version", handler.Version)
	router.POST("/update", handler.Update)
}
