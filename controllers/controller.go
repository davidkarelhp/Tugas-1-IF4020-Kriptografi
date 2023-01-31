package controllers

import (
	"Tugas-1-IF4020-Kriptografi/json"
	"Tugas-1-IF4020-Kriptografi/services"
	"fmt"
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
	PostHill(ctx *gin.Context)
}

type Controller struct {
	service services.IService
}

//NewController is creating anew instance of Controlller
func NewController() IController {
	return &Controller{
		service: services.NewService(),
	}
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

func (c *Controller) PostHill(ctx *gin.Context) {
	var req json.HillReq
	if err := ctx.BindJSON(&req); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect format",
			"success": false,
		})
		return
	}

	encInt, err := req.Encrypt.Int64()
	if err != nil || (encInt != 0 && encInt != 1) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Encrypt should be an integer (0 or 1)",
			"success": false,
		})
		return
	}

	mInt, err := req.M.Int64()
	if err != nil || mInt <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "M should be an integer (> 0)",
			"success": false,
		})
		return
	}

	typeInt, err := req.Type.Int64()
	if err != nil || (typeInt != 0 && typeInt != 1) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Type has incorrect format",
			"success": false,
		})
		return
	}

	var encrypt bool
	if encInt == 1 {
		encrypt = true
	} else {
		encrypt = false
	}

	var result string
	if typeInt == 0 {
		result, err = c.service.HillCipher(req.InputText, req.Key, int(mInt), encrypt)
	} else {

	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Encrypt or Decrypt successful",
		"success": true,
		"result":  result,
	})
}
