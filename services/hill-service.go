package services

import (
	"fmt"
	"math"
	"mime/multipart"
	"strings"

	"gonum.org/v1/gonum/mat"
)

type IHillService interface {
	HillCipher(textString string, matrixString string, m int, encrypt bool) (string, error)
	HillCipherFile(textFileHeader *multipart.FileHeader, matrixString string, m int, encrypt bool) (string, error)
}

type HillService struct {
	cs ICommonService
}

//NewController is creating anew instance of Controlller
func NewHillService() IHillService {
	return &HillService{
		cs: NewCommonService(),
	}
}

func (src *HillService) HillCipherFile(textFileHeader *multipart.FileHeader, matrixString string, m int, encrypt bool) (string, error) {
	textString, err := src.cs.ReadTxtFile(textFileHeader)
	if err != nil {
		return "", nil
	}

	return src.HillCipher(textString, matrixString, m, encrypt)
}

func (src *HillService) HillCipher(textString string, matrixString string, m int, encrypt bool) (string, error) {
	ASCII_OFFSET := 65
	numOfSymbols := 26
	matrix, err := src.cs.ParseStringToMatrix(matrixString, m)

	if err != nil {
		return "", err
	}

	logDet, sign := mat.LogDet(matrix)
	det := sign * math.Exp(logDet)

	var ret string
	var matInv mat.Dense
	var matInvScaled mat.Dense

	err = matInv.Inverse(matrix)
	if err != nil {
		return "", NewCustomError("Matrix doesn't have an inverse")
	}

	detMod := int(det) % numOfSymbols
	if detMod < 0 {
		detMod = numOfSymbols + detMod
	}
	invMod := src.cs.ModInverse(detMod, numOfSymbols)
	fmt.Println("invMod = ", invMod)

	if invMod == -1 {
		return "", NewCustomError("There is no modular multiplicative inverse")
	}

	matInvScaled.Scale(det*float64(invMod), &matInv)
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			modMatElement := int(math.Round(matInvScaled.At(i, j))) % numOfSymbols
			if modMatElement < 0 {
				modMatElement = numOfSymbols + modMatElement
			}
			matInvScaled.Set(i, j, float64(modMatElement))
		}
	}

	textString = strings.ToUpper(textString)
	runes := []rune(textString)
	runes = src.cs.FilterRunesAZ(runes)

	if len(runes)%m != 0 {
		return "", NewCustomError("Input text length should be the multiple of M")
	}

	chunks := src.cs.ChunkSlice(runes, m, ASCII_OFFSET)
	for _, chunk := range chunks {
		pmat := mat.NewDense(m, 1, chunk)
		var cmat mat.Dense
		if encrypt {
			cmat.Mul(matrix, pmat)
		} else {
			cmat.Mul(&matInvScaled, pmat)
		}

		for i := 0; i < m; i++ {
			f := math.Round(cmat.At(i, 0))
			symbolASCII := (int(f) % numOfSymbols)
			if symbolASCII < 0 {
				symbolASCII = symbolASCII + numOfSymbols
			}
			ret += string(rune(symbolASCII + ASCII_OFFSET))
		}
	}
	return ret, nil
}
