package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kazuki0924/go-gin/controller"
	"github.com/kazuki0924/go-gin/middlewares"
	"github.com/kazuki0924/go-gin/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	// server := gin.Default()
	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger())

	// server.Use(gin.Logger())

	// server.GET("/test", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "OK!!",
	// 	})
	// })

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}
