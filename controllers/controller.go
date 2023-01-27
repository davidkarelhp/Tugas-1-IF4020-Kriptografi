package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IController interface {
	HelloWorld(ctx *gin.Context)
	GGWP(ctx *gin.Context)
	Index(ctx *gin.Context)
}

type Controller struct {
}

//NewController is creating anew instance of Controlller
func NewController() IController {
	return &Controller{}
}

func (c *Controller) HelloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func (c *Controller) GGWP(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home/ggwp.html", gin.H{
		"title": "GGWP",
	})
}

func (c *Controller) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home/index.html", gin.H{
		"title": "Home",
	})
}
