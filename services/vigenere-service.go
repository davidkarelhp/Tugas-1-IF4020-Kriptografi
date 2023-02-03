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
	bytes, err := src.cs.ReadFileBytes(textFileHeader)
	if err != nil {
		return "", err
	}

	textString := string(bytes)

	res, err := src.VigenereCipher(textString, key, encrypt)
	if err != nil {
		fmt.Println("ERROR = ", err.Error())
		return "", err
	}

	return res, nil
}

func (src *VigenereService) VigenereCipher(textString string, key string, encrypt bool) (string, error) {
	res := []rune{}
	// char := ""
	var char rune
	j := 0

	key = strings.ToUpper(key)
	keyRunes := []rune(key)
	keyRunes = src.cs.FilterRunesAZ(keyRunes)

	textString = strings.ToUpper(textString)
	textRunes := []rune(textString)
	textRunes = src.cs.FilterRunesAZ(textRunes)
	keyLen := len(keyRunes)

	for i := 0; i < len(textRunes); i++ {
		if encrypt {
			p := textRunes[i] - 65
			k := keyRunes[j] - 65
			char = rune(((p + k) % 26) + 65)

		} else {
			p := textRunes[i] - 65
			k := keyRunes[j] - 65
			char = rune(src.cs.ModNegatif(int(p-k), 26) + 65)
		}
		res = append(res, char)
		j++
		if j == keyLen {
			j = 0
		}
	}
	return string(res), nil
}
