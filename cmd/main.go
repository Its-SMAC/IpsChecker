package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("web/templates/*")
	r.Static("/static", "./web/static")

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "Api is fine"})
	})

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{"Titulo": "Ip Checker"})
	})

	fmt.Println("Server has been start.")
	fmt.Println("Server run on 127.0.0.1:8080")

	r.Run(":8080")
}
