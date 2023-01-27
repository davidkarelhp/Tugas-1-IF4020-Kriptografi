package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IController interface {
	HelloWorld(ctx *gin.Context)
	GGWP(ctx *gin.Context)
	Index(ctx *gin.Context)
	Vignere(ctx *gin.Context)
	AutoKeyVignere(ctx *gin.Context)
	ExtendedVignere(ctx *gin.Context)
	Affine(ctx *gin.Context)
	Playfair(ctx *gin.Context)
	Hill(ctx *gin.Context)
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

func (c *Controller) Vignere(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "encrypt-decrypt/vignere.html", gin.H{
		"title": "Vignere Cipher",
	})
}

func (c *Controller) AutoKeyVignere(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "encrypt-decrypt/auto-key-vignere.html", gin.H{
		"title": "Auto-key Vignere Cipher",
	})
}

func (c *Controller) ExtendedVignere(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "encrypt-decrypt/extended-vignere.html", gin.H{
		"title": "Extended Vignere Cipher",
	})
}

func (c *Controller) Affine(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "encrypt-decrypt/affine.html", gin.H{
		"title": "Affine Cipher",
	})
}

func (c *Controller) Playfair(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "encrypt-decrypt/playfair.html", gin.H{
		"title": "Playfair Cipher",
	})
}

func (c *Controller) Hill(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "encrypt-decrypt/hill.html", gin.H{
		"title": "Hill Cipher",
	})
}
