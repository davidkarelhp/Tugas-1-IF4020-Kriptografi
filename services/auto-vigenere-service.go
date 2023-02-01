package services

import (
	"fmt"
	// "math"
	"mime/multipart"
	"strings"

	// "gonum.org/v1/gonum/mat"
)

type IAutoVigenereService interface {
	AutoVigenereCipher(textString string, key string, encrypt bool) (string, error)
	AutoVigenereCipherFile(textFileHeader *multipart.FileHeader, key string, encrypt bool) (string, error)
}

type AutoVigenereService struct {
	cs ICommonService
}

//NewAutoVigenereService is creating a new instance of AutoVigenereService
func NewAutoVigenereService() IAutoVigenereService {
	return &AutoVigenereService{
		cs: NewCommonService(),
	}
}

func (src *AutoVigenereService) AutoVigenereCipherFile(textFileHeader *multipart.FileHeader, key string, encrypt bool) (string, error) {
	textString, err := src.cs.ReadTxtFile(textFileHeader)
	if err != nil {
		return "", err
	}

	res, err := src.AutoVigenereCipher(textString, key, encrypt)
	if err != nil {
		fmt.Println("ERROR = ", err.Error())
		return "", err
	}

	return res, nil
}

func (src *AutoVigenereService) AutoVigenereCipher (textString string, key string, encrypt bool) (string, error) {
	res := ""
	char := ""
	j := 0 

	key = strings.ToUpper(key)
	keyRunes := []rune(key)
	keyRunes = src.cs.FilterRunesAZ(keyRunes)
	keyRunes = src.cs.RemoveRune(keyRunes, rune(74))

	textString = strings.ToUpper(textString)
	textRunes := []rune(textString)
	textRunes = src.cs.FilterRunesAZ(textRunes)
	textRunes = src.cs.ReplaceRune(textRunes, rune(74), rune(73))

	for i := 0; i < len(textRunes); i++ {
		if encrypt{
			p := textRunes[i] - 65
			k := keyRunes[i] - 65
			char = string(((p + k) % 26) + 65)
			
		} else {
			p := textRunes[i] - 65
			k := keyRunes[i] - 65
			char = string(rune(src.cs.ModLikePython(int(p-k), 26) + 65))
		}
		res = res + char
		j++
		if j == len(key) {
			j = 0
		}
	}
	return res, nil
}