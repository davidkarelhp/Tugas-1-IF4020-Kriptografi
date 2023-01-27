package main

import (
	"Tugas-1-IF4020-Kriptografi/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*.html")

	var (
		controller controllers.IController = controllers.NewController()
	)

	fmt.Println("Hello World!")

	r.GET("/hello-world", func(c *gin.Context) { controller.HelloWorld(c) })

	r.GET("/ggwp", func(c *gin.Context) { controller.GGWP(c) })

	r.GET("/", func(c *gin.Context) { controller.Index(c) })

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
