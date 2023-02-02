package services

import (
	"fmt"
	"mime/multipart"
)

type IEVigenereService interface {
	EVigenereCipher(textString string, key string, encrypt bool) (string, error)
	EVigenereCipherFile(textFileHeader *multipart.FileHeader, key string, encrypt bool) (string, error)
}

type EVigenereService struct {
	cs ICommonService
}

//NewEVigenereService is creating a new instance of EVigenereService
func NewEVigenereService() IEVigenereService {
	return &EVigenereService{
		cs: NewCommonService(),
	}
}

func (src *EVigenereService) EVigenereCipherFile(textFileHeader *multipart.FileHeader, key string, encrypt bool) (string, error) {
	textString, err := src.cs.ReadTxtFile(textFileHeader)
	if err != nil {
		return "", err
	}

	res, err := src.EVigenereCipher(textString, key, encrypt)
	if err != nil {
		fmt.Println("ERROR = ", err.Error())
		return "", err
	}

	return res, nil
}

func (src *EVigenereService) EVigenereCipher(textString string, key string, encrypt bool) (string, error) {
	res := ""
	char := ""
	j := 0

	keyRunes := []rune(key)
	textRunes := []rune(textString)

	keyLen := len(keyRunes)

	for i := 0; i < len(textRunes); i++ {
		if encrypt {
			char = string((textRunes[i] + keyRunes[j]) % 256)

		} else {
			char = string(rune(src.cs.ModLikePython(int(textRunes[i]-keyRunes[j]), 256)))
		}
		res = res + char
		j++
		if j == keyLen {
			j = 0
		}
	}
	return res, nil
}
