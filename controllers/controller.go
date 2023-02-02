package controllers

import (
	"Tugas-1-IF4020-Kriptografi/binding_struct"
	"Tugas-1-IF4020-Kriptografi/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IController interface {
	HelloWorld(ctx *gin.Context)
	GGWP(ctx *gin.Context)
	Index(ctx *gin.Context)
	Vigenere(ctx *gin.Context)
	AutoKeyVigenere(ctx *gin.Context)
	ExtendedVigenere(ctx *gin.Context)
	Affine(ctx *gin.Context)
	Playfair(ctx *gin.Context)
	Hill(ctx *gin.Context)
	PostHill(ctx *gin.Context)
	PostPlayfair(ctx *gin.Context)
	PostVigenere(ctx *gin.Context)
	PostAutoKeyVigenere(ctx *gin.Context)
	PostExtendedVigenere(ctx *gin.Context)
	PostAffine(ctx *gin.Context)
}

type Controller struct {
	hs  services.IHillService
	ps  services.IPlayfairService
	vs  services.IVigenereService
	avs services.IAutoVigenereService
	evs services.IEVigenereService
	afs services.IAffineService
}

//NewController is creating anew instance of Controlller
func NewController() IController {
	return &Controller{
		hs:  services.NewHillService(),
		ps:  services.NewPlayfairService(),
		vs:  services.NewVigenereService(),
		avs: services.NewAutoVigenereService(),
		evs: services.NewEVigenereService(),
		afs: services.NewAffineService(),
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

func (c *Controller) Vigenere(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "encrypt-decrypt/vigenere.html", gin.H{
		"title": "Vigenere Cipher",
	})
}

func (c *Controller) AutoKeyVigenere(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "encrypt-decrypt/auto-key-vigenere.html", gin.H{
		"title": "Auto-key Vigenere Cipher",
	})
}

func (c *Controller) ExtendedVigenere(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "encrypt-decrypt/extended-vigenere.html", gin.H{
		"title": "Extended Vigenere Cipher",
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
	var req binding_struct.HillReq
	if err := ctx.ShouldBind(&req); err != nil {
		fmt.Println("ERROR: ", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect format",
			"success": false,
		})
		return
	}

	encInt := req.Encrypt
	mInt := req.M
	typeInt := req.Type

	var encrypt bool
	if encInt == 1 {
		encrypt = true
	} else {
		encrypt = false
	}

	var result string
	var err error

	if typeInt == 0 {
		result, err = c.hs.HillCipher(req.InputText, req.Key, int(mInt), encrypt)
	} else {
		file, fileErr := ctx.FormFile("file")
		if fileErr != nil {
			fmt.Println("ERROR: ", fileErr.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Incorrect format",
				"success": false,
			})
			return
		}

		result, err = c.hs.HillCipherFile(file, req.Key, int(mInt), encrypt)
	}

	if err != nil {
		fmt.Println("ERROR: ", err.Error())
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

func (c *Controller) PostPlayfair(ctx *gin.Context) {
	var req binding_struct.PlayfairReq
	if err := ctx.ShouldBind(&req); err != nil {
		fmt.Println("ERROR: ", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect format",
			"success": false,
		})
		return
	}

	encInt := req.Encrypt
	typeInt := req.Type

	var encrypt bool
	if encInt == 1 {
		encrypt = true
	} else {
		encrypt = false
	}

	var result string
	var err error

	if typeInt == 0 {
		result, err = c.ps.PlayfairCipher(req.InputText, req.Key, encrypt)
	} else {
		file, fileErr := ctx.FormFile("file")
		if fileErr != nil {
			fmt.Println("ERROR: ", fileErr.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Incorrect format",
				"success": false,
			})
			return
		}

		result, err = c.ps.PlayfairCipherFile(file, req.Key, encrypt)
	}

	if err != nil {
		fmt.Println("ERROR: ", err.Error())
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

func (c *Controller) PostAutoKeyVigenere(ctx *gin.Context) {
	var req binding_struct.VigenereReq
	if err := ctx.ShouldBind(&req); err != nil {
		fmt.Println("ERROR: ", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect format",
			"success": false,
		})
		return
	}

	encInt := req.Encrypt
	typeInt := req.Type

	var encrypt bool
	if encInt == 1 {
		encrypt = true
	} else {
		encrypt = false
	}

	var result string
	var err error

	if typeInt == 0 {
		result, err = c.avs.AutoVigenereCipher(req.InputText, req.Key, encrypt)
	} else {
		file, fileErr := ctx.FormFile("file")
		if fileErr != nil {
			fmt.Println("ERROR: ", fileErr.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Incorrect format",
				"success": false,
			})
			return
		}

		result, err = c.avs.AutoVigenereCipherFile(file, req.Key, encrypt)
	}

	if err != nil {
		fmt.Println("ERROR: ", err.Error())
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

func (c *Controller) PostVigenere(ctx *gin.Context) {
	var req binding_struct.VigenereReq
	if err := ctx.ShouldBind(&req); err != nil {
		fmt.Println("ERROR: ", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect format",
			"success": false,
		})
		return
	}

	encInt := req.Encrypt
	typeInt := req.Type

	var encrypt bool
	if encInt == 1 {
		encrypt = true
	} else {
		encrypt = false
	}

	var result string
	var err error

	if typeInt == 0 {
		result, err = c.vs.VigenereCipher(req.InputText, req.Key, encrypt)
	} else {
		file, fileErr := ctx.FormFile("file")
		if fileErr != nil {
			fmt.Println("ERROR: ", fileErr.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Incorrect format",
				"success": false,
			})
			return
		}

		result, err = c.vs.VigenereCipherFile(file, req.Key, encrypt)
	}

	if err != nil {
		fmt.Println("ERROR: ", err.Error())
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

func (c *Controller) PostExtendedVigenere(ctx *gin.Context) {
	var req binding_struct.VigenereReq
	if err := ctx.ShouldBind(&req); err != nil {
		fmt.Println("ERROR: ", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect format",
			"success": false,
		})
		return
	}

	encInt := req.Encrypt
	typeInt := req.Type

	var encrypt bool
	if encInt == 1 {
		encrypt = true
	} else {
		encrypt = false
	}

	var result string
	var err error

	if typeInt == 0 {
		result, err = c.evs.EVigenereCipher(req.InputText, req.Key, encrypt)
	} else {
		file, fileErr := ctx.FormFile("file")
		if fileErr != nil {
			fmt.Println("ERROR: ", fileErr.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Incorrect format",
				"success": false,
			})
			return
		}

		bytes, errBytes := c.evs.EVigenereCipherFile(file, req.Key, encrypt)
		if errBytes == nil {
			fileName := file.Filename
			ctx.Header("Content-Disposition", "attachment; filename="+fileName)
			ctx.Header("Content-Type", "application/octet-stream")
			ctx.Header("Accept-Length", fmt.Sprintf("%d", len(bytes)))
			ctx.Writer.Write(bytes)
			return

		}
		err = errBytes
	}

	if err != nil {
		fmt.Println("ERROR: ", err.Error())
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

func (c *Controller) PostAffine(ctx *gin.Context) {
	var req binding_struct.AffineReq
	if err := ctx.ShouldBind(&req); err != nil {
		fmt.Println("ERROR: ", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect format",
			"success": false,
		})
		return
	}

	encInt := req.Encrypt
	typeInt := req.Type

	var encrypt bool
	if encInt == 1 {
		encrypt = true
	} else {
		encrypt = false
	}

	var result string
	var err error

	if typeInt == 0 {
		result, err = c.afs.AffineCipher(req.InputText, req.M, req.B, encrypt)
	} else {
		file, fileErr := ctx.FormFile("file")
		if fileErr != nil {
			fmt.Println("ERROR: ", fileErr.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Incorrect format",
				"success": false,
			})
			return
		}

		result, err = c.afs.AffineCipherFile(file, req.M, req.B, encrypt)
	}

	if err != nil {
		fmt.Println("ERROR: ", err.Error())
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
