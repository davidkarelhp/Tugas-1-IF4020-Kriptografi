package services

import (
	"fmt"
	"mime/multipart"
	"strings"
)

type IVigenereService interface {
	VigenereCipher(textString string, key string, encrypt bool) (string, error)
	VigenereCipherFile(textFileHeader *multipart.FileHeader, key string, encrypt bool) (string, error)
}

type VigenereService struct {
	cs ICommonService
}

//NewVigenereService is creating a new instance of VigenereService
func NewVigenereService() IVigenereService {
	return &VigenereService{
		cs: NewCommonService(),
	}
}

func (src *VigenereService) VigenereCipherFile(textFileHeader *multipart.FileHeader, key string, encrypt bool) (string, error) {
	textString, err := src.cs.ReadTxtFile(textFileHeader)
	if err != nil {
		return "", err
	}

	res, err := src.VigenereCipher(textString, key, encrypt)
	if err != nil {
		fmt.Println("ERROR = ", err.Error())
		return "", err
	}

	return res, nil
}

func (src *VigenereService) VigenereCipher (textString string, key string, encrypt bool) (string, error) {
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