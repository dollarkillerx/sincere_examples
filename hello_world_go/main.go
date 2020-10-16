package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()

	app.GET("/hello", func(context *gin.Context) {
		context.String(200,"Hello World")
	})

	app.Run("0.0.0.0:8081")
}
