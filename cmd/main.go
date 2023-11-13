package main

import (
	"gdrive_uploader/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	api := server.Group("/api")
	api.GET("/auth", handler.Auth)
	api.GET("/callback", handler.Callback)
	api.GET("/file", handler.File)
	api.POST("/upload", handler.Upload)

	server.Run("localhost:8081")
}
