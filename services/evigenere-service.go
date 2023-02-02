package services

import (
	"fmt"
	"mime/multipart"
)

type IEVigenereService interface {
	EVigenereCipher(textString string, key string, encrypt bool) (string, error)
	EVigenereCipherFile(textFileHeader *multipart.FileHeader, key string, encrypt bool) ([]byte, error)
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

func (src *EVigenereService) EVigenereCipherFile(textFileHeader *multipart.FileHeader, key string, encrypt bool) ([]byte, error) {
	bytes, err := src.cs.ReadFileBytes(textFileHeader)
	if err != nil {
		return nil, err
	}

	res, err := src.EVigenereCipherBytes(bytes, key, encrypt)
	if err != nil {
		fmt.Println("ERROR = ", err.Error())
		return nil, err
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

func (src *EVigenereService) EVigenereCipherBytes(bytes []byte, key string, encrypt bool) ([]byte, error) {
	res := []byte{}
	var singleByte byte
	j := 0

	keyRunes := []rune(key)

	keyLen := len(keyRunes)

	for i := 0; i < len(bytes); i++ {
		if encrypt {
			singleByte = byte((int(bytes[i]) + int(keyRunes[j])) % 256)

		} else {
			singleByte = byte(rune(src.cs.ModLikePython(int(bytes[i])-int(keyRunes[j]), 256)))
		}
		res = append(res, singleByte)
		j++
		if j == keyLen {
			j = 0
		}

	}

	return res, nil
}
