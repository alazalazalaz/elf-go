package route

import (
	"elf-go/example/app/http/handler"
	"elf-go/example/app/http/middleware"
	middleware2 "elf-go/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoute(router *gin.Engine) {
	router.Use(gin.Recovery())
	router.Use(middleware.Cors, middleware2.PrintReqAndResp)

	router.GET("/info", handler.Info)
	router.GET("/get-user-info", handler.GetUserInfo)
	router.GET("/version", handler.Version)
	router.POST("/update", handler.Update)
	router.GET("/panic", handler.Panic)
}
