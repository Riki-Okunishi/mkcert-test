package main

import (
	"io"
	"os"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file)
	log.SetOutput(gin.DefaultWriter)

	router := gin.Default()

	router.GET("/proxy", myMiddleware(), func(context *gin.Context) {
		context.JSON(200, gin.H{
			"key": "value",
		})
	})

	router.RunTLS(":6121", "./cert/localhost.pem", "./cert/localhost-key.pem")
}

func myMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		header := context.GetHeader("X-REQUEST-ID")
		log.Print("RequestId=" + header)
		context.Next()
	}
}
