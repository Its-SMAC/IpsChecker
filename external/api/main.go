package main

import (
	"fmt"
	"html/template"
	"io/fs"
	"ipchecker/internal"
	"ipchecker/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	webFS := web.GetEmbed()

	tmpl, err := template.ParseFS(webFS, "templates/*")
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(tmpl)

	staticFS, err := fs.Sub(webFS, "static")
	if err != nil {
		panic(err)
	}

	r.StaticFS("/static", http.FS(staticFS))

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "Api is fine"})
	})

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{"Titulo": "Ip Checker"})
	})

	r.POST("/scan/ip", func(ctx *gin.Context) {
		var request struct {
			Alvo string `json:"alvo"`
		}

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
			return
		}

		lista := internal.Check(request.Alvo)

		ctx.JSON(http.StatusOK, lista)
	})

	fmt.Println("Server has been start.")
	fmt.Println("Server run on 127.0.0.1:8080")

	r.Run(":8080")
}
