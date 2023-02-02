package services

import (
	"fmt"
	// "math"
	"mime/multipart"
	"strings"
	// "gonum.org/v1/gonum/mat"
)

type IAffineService interface {
	AffineCipher(textString string, key int, encrypt bool) (string, error)
	AffineCipherFile(textFileHeader *multipart.FileHeader, key int, encrypt bool) (string, error)
}

type AffineService struct {
	cs ICommonService
}

//NewAffineService is creating a new instance of AffineService
func NewAffineService() IAffineService {
	return &AffineService{
		cs: NewCommonService(),
	}
}

func (src *AffineService) AffineCipherFile(textFileHeader *multipart.FileHeader, key int, encrypt bool) (string, error) {
	textString, err := src.cs.ReadTxtFile(textFileHeader)
	if err != nil {
		return "", err
	}

	res, err := src.AffineCipher(textString, key, encrypt)
	if err != nil {
		fmt.Println("ERROR = ", err.Error())
		return "", err
	}

	return res, nil
}

func (src *AffineService) AffineCipher(textString string, key int, encrypt bool) (string, error) {
	res := ""
	char := ""

	textString = strings.ToUpper(textString)
	textRunes := []rune(textString)
	textRunes = src.cs.FilterRunesAZ(textRunes)
	textRunes = src.cs.ReplaceRune(textRunes, rune(74), rune(73))
	m := 0

	for i := 0; i < len(textRunes); i++ {
		if encrypt {
			p := textRunes[i] - 65
			char = string(((m*int(p) + key) % 26) + 65) //gw mikirny ini m masukan dropdown aja

		} else {
			p := textRunes[i] - 65
			m = src.cs.ModInverse(m, 26)
			char = string(rune(src.cs.ModLikePython(m*(int(p)-key), 26) + 65))
		}
		res = res + char
	}
	return res, nil
}
