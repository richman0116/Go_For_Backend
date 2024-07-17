package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	g := gin.Default()

	g.GET("/message", func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello from Go backend",
		})
	})

	g.ServeHTTP(w, r)
}

func main() {
	LoadAppConfig()

	port := fmt.Sprintf(":%s", AppConfig.Port)
	http.HandleFunc("/", Handler)
	http.ListenAndServe(port, nil)
}
