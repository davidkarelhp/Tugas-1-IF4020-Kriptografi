package services

import (
	"fmt"
	"mime/multipart"
	"strings"
)

type IPlayfairService interface {
	PlayfairCipher(textString string, key string, encrypt bool) (string, error)
	PlayfairCipherFile(textFileHeader *multipart.FileHeader, key string, encrypt bool) (string, error)
}

type PlayfairService struct {
	cs ICommonService
}

//NewPlayfairService is creating a new instance of PlayfairService
func NewPlayfairService() IPlayfairService {
	return &PlayfairService{
		cs: NewCommonService(),
	}
}

func (src *PlayfairService) PlayfairCipherFile(textFileHeader *multipart.FileHeader, key string, encrypt bool) (string, error) {
	textString, err := src.cs.ReadTxtFile(textFileHeader)
	if err != nil {
		return "", err
	}

	res, err := src.PlayfairCipher(textString, key, encrypt)
	if err != nil {
		fmt.Println("ERROR = ", err.Error())
		return "", err
	}

	return res, nil
}

func (src *PlayfairService) PlayfairCipher(textString string, key string, encrypt bool) (string, error) {
	key = strings.ToUpper(key)
	keyRunes := []rune(key)
	keyRunes = src.cs.FilterRunesAZ(keyRunes)
	keyRunes = src.cs.FilterDuplicateValues(keyRunes)
	keyRunes = src.cs.RemoveRune(keyRunes, rune(74))

	matrix := make([][]rune, 5)
	matrixKey := make(map[rune]*IntPair)
	for i := range matrix {
		matrix[i] = make([]rune, 5)
	}

	matrix, matrixKey = src.generatePlayfairMatrix(keyRunes)
	fmt.Println(matrix)

	textString = strings.ToUpper(textString)
	textRunes := []rune(textString)
	textRunes = src.cs.FilterRunesAZ(textRunes)
	textRunes = src.cs.ReplaceRune(textRunes, rune(74), rune(73))

	shifter := -1
	if encrypt {
		shifter = 1
	}

	ret := ""

	bigrams := src.splitRunesToBigram(textRunes)
	for _, bigram := range bigrams {
		if matrixKey[rune(bigram.first)].first == matrixKey[rune(bigram.second)].first {
			i := matrixKey[rune(bigram.first)].first
			newFirstJ := (matrixKey[rune(bigram.first)].second + shifter) % 5
			if newFirstJ < 0 {
				newFirstJ += 5
			}

			newSecondJ := (matrixKey[rune(bigram.second)].second + shifter) % 5
			if newSecondJ < 0 {
				newSecondJ += 5
			}

			ret += string(matrix[i][newFirstJ] + 65)
			ret += string(matrix[i][newSecondJ] + 65)

		} else if matrixKey[rune(bigram.first)].second == matrixKey[rune(bigram.second)].second {
			j := matrixKey[rune(bigram.first)].second
			newFirstI := (matrixKey[rune(bigram.first)].first + shifter) % 5
			if newFirstI < 0 {
				newFirstI += 5
			}

			newSecondI := (matrixKey[rune(bigram.second)].first + shifter) % 5
			if newSecondI < 0 {
				newSecondI += 5
			}

			ret += string(matrix[newFirstI][j] + 65)
			ret += string(matrix[newSecondI][j] + 65)

		} else {
			ret += string(matrix[matrixKey[rune(bigram.first)].first][matrixKey[rune(bigram.second)].second] + 65)
			ret += string(matrix[matrixKey[rune(bigram.second)].first][matrixKey[rune(bigram.first)].second] + 65)
		}
	}

	return ret, nil
}

func (src *PlayfairService) generatePlayfairMatrix(keyRunes []rune) ([][]rune, map[rune]*IntPair) {
	retMatrix := make([][]rune, 5)
	retMatrixKey := make(map[rune]*IntPair)
	for i := range retMatrix {
		retMatrix[i] = make([]rune, 5)
	}

	keyIterator := 0
	alphabetIterator := rune(0)
	runeLen := len(keyRunes)

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i*5+j < runeLen {
				retMatrix[i][j] = keyRunes[keyIterator] - 65
				retMatrixKey[keyRunes[keyIterator]-65] = &IntPair{i, j}
				keyIterator++
			} else {
				fmt.Println("ALPHABET_ITERATOR", alphabetIterator)
				for alphabet := alphabetIterator; alphabet < 26; alphabet++ {
					if alphabet != 74-65 {
						_, ok := retMatrixKey[alphabet]
						if !ok {
							retMatrix[i][j] = alphabet
							retMatrixKey[alphabet] = &IntPair{i, j}
							alphabetIterator = alphabet + 1
							break
						}
					}
				}

			}
			// fmt.Println(retMatrix)
		}
	}
	// fmt.Println(retMatrix)
	return retMatrix, retMatrixKey
}

func (src *PlayfairService) splitRunesToBigram(runes []rune) []*IntPair {
	ret := []*IntPair{}
	newBigram := true
	retIterator := 0
	for i := 0; i < len(runes); i++ {
		if newBigram {
			ret = append(ret, &IntPair{int(runes[i]) - 65, 0})
			newBigram = false
		} else {
			if rune(ret[retIterator].first) == runes[i]-65 {
				ret[retIterator].second = 88 - 65
				i--
			} else {
				ret[retIterator].second = int(runes[i]) - 65
			}
			retIterator++
			newBigram = true
		}
	}

	if !newBigram {
		ret[retIterator].second = 88 - 65
	}

	return ret
}
