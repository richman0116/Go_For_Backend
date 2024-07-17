package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	LoadAppConfig()

	r := gin.Default()

	r.GET("/message", func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello from Go backend",
		})
	})

	r.Run(fmt.Sprintf(":%s", AppConfig.Port))
}
