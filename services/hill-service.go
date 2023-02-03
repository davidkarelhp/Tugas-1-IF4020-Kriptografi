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

//NewHillService is creating a new instance of HillService
func NewHillService() IHillService {
	return &HillService{
		cs: NewCommonService(),
	}
}

func (src *HillService) HillCipherFile(textFileHeader *multipart.FileHeader, matrixString string, m int, encrypt bool) (string, error) {
	bytes, err := src.cs.ReadFileBytes(textFileHeader)
	if err != nil {
		return "", err
	}

	textString := string(bytes)

	res, err := src.HillCipher(textString, matrixString, m, encrypt)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return res, nil
}

func (src *HillService) HillCipher(textString string, matrixString string, m int, encrypt bool) (string, error) {
	asciiOffset := 65
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
	// fmt.Println("invMod = ", invMod)

	if invMod == -1 {
		return "", NewCustomError("There is no modular multiplicative inverse of (determinant of matrix (key) % 26) under modulo 26")
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

	chunks := src.cs.ChunkSlice(runes, m, asciiOffset)
	for _, chunk := range chunks {
		imat := mat.NewDense(m, 1, chunk)
		var omat mat.Dense
		if encrypt {
			omat.Mul(matrix, imat)
		} else {
			omat.Mul(&matInvScaled, imat)
		}

		for i := 0; i < m; i++ {
			f := math.Round(omat.At(i, 0))
			symbolASCII := (int(f) % numOfSymbols)
			if symbolASCII < 0 {
				symbolASCII = symbolASCII + numOfSymbols
			}
			ret += string(rune(symbolASCII + asciiOffset))
		}
	}
	return ret, nil
}
