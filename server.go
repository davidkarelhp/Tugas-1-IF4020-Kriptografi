package main

import (
	"Tugas-1-IF4020-Kriptografi/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*.html")
	r.Static("/assets", "./assets")

	var (
		controller controllers.IController = controllers.NewController()
	)

	fmt.Println("Hello World!")

	r.GET("/hello-world", func(c *gin.Context) { controller.HelloWorld(c) })

	r.GET("/ggwp", func(c *gin.Context) { controller.GGWP(c) })

	r.GET("/", func(c *gin.Context) { controller.Index(c) })

	r.GET("/vignere", func(c *gin.Context) { controller.Vignere(c) })

	r.GET("/auto-key-vignere", func(c *gin.Context) { controller.AutoKeyVignere(c) })

	r.GET("/extended-vignere", func(c *gin.Context) { controller.ExtendedVignere(c) })

	r.GET("/affine", func(c *gin.Context) { controller.Affine(c) })

	r.GET("/playfair", func(c *gin.Context) { controller.Playfair(c) })

	r.GET("/hill", func(c *gin.Context) { controller.Hill(c) })
	r.POST("/hill", func(c *gin.Context) { controller.PostHill(c) })

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
