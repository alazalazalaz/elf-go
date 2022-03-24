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

	router.GET("/", handler.Info)
	router.GET("/info", handler.Info)
	router.GET("/slow-query", handler.SlowQuery)
	router.GET("/version", handler.Version)
	router.GET("/panic", handler.Panic)
	router.GET("/login", handler.Login)
	router.POST("/sleep", handler.Sleep)
	router.GET("/loop", handler.Loop)

	authRouter := router.Group("auth")
	authRouter.Use(middleware2.ParseJwt)
	authRouter.GET("", handler.Auth)

	// select
	router.GET("/select", handler.Select)
	// create
	router.GET("/create", handler.Create)
	// update
	router.GET("/update", handler.Update)
	// save
	router.GET("/save", handler.Save)

	restfulController := new(handler.RestfulController)
	restfulControllerR := router.Group("restful")
	{
		restfulControllerR.GET("/get", restfulController.Get)
		restfulControllerR.PUT("/put", restfulController.Put)
		restfulControllerR.POST("/post", restfulController.Post)
		restfulControllerR.DELETE("/delete", restfulController.Delete)
		restfulControllerR.POST("/file", restfulController.File)
	}

}
