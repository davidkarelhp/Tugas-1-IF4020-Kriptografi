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

	r.GET("/vigenere", func(c *gin.Context) { controller.Vigenere(c) })
	r.POST("/vigenere", func(c *gin.Context) { controller.PostVigenere(c) })

	r.GET("/auto-key-vigenere", func(c *gin.Context) { controller.AutoKeyVigenere(c) })
	r.POST("/auto-key-vigenere", func(c *gin.Context) { controller.PostAutoKeyVigenere(c) })

	r.GET("/extended-vigenere", func(c *gin.Context) { controller.ExtendedVigenere(c) })
	r.POST("/extended-vigenere", func(c *gin.Context) { controller.PostExtendedVigenere(c) })

	r.GET("/affine", func(c *gin.Context) { controller.Affine(c) })
	r.POST("/affine", func(c *gin.Context) { controller.PostAffine(c) })

	r.GET("/playfair", func(c *gin.Context) { controller.Playfair(c) })
	r.POST("/playfair", func(c *gin.Context) { controller.PostPlayfair(c) })

	r.GET("/hill", func(c *gin.Context) { controller.Hill(c) })
	r.POST("/hill", func(c *gin.Context) { controller.PostHill(c) })

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
