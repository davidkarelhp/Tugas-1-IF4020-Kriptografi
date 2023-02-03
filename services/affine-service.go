package services

import (
	"fmt"
	// "math"
	"mime/multipart"
	"strings"
	// "gonum.org/v1/gonum/mat"
)

type IAffineService interface {
	AffineCipher(textString string, m int, b int, encrypt bool) (string, error)
	AffineCipherFile(textFileHeader *multipart.FileHeader, m int, b int, encrypt bool) (string, error)
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

func (src *AffineService) AffineCipherFile(textFileHeader *multipart.FileHeader, m int, b int, encrypt bool) (string, error) {
	textString, err := src.cs.ReadTxtFile(textFileHeader)
	if err != nil {
		return "", err
	}

	res, err := src.AffineCipher(textString, m, b, encrypt)
	if err != nil {
		fmt.Println("ERROR = ", err.Error())
		return "", err
	}

	return res, nil
}

func (src *AffineService) AffineCipher(textString string, m int, b int, encrypt bool) (string, error) {
	var gcd int
	if m > 26 {
		gcd = src.cs.GCD(m, 26)
	} else {
		gcd = src.cs.GCD(26, m)
	}

	if gcd != 1 {
		return "", NewCustomError("M is not relatively prime with 26")
	}

	if b < 1 || b > 25 {
		return "", NewCustomError("B is either smaller than 1 or greater than 25")
	}

	res := ""
	char := ""

	textString = strings.ToUpper(textString)
	textRunes := []rune(textString)
	textRunes = src.cs.FilterRunesAZ(textRunes)

	if !encrypt {
		m = src.cs.ModInverse(m, 26)
	}

	for i := 0; i < len(textRunes); i++ {
		if encrypt {
			p := textRunes[i] - 65
			char = string(rune(((m*int(p) + b) % 26) + 65))
		} else {
			p := textRunes[i] - 65
			char = string(rune(src.cs.ModNegatif(m*(int(p)-b), 26) + 65))
		}
		res = res + char
	}
	return res, nil
}
