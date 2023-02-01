package services

import (
	"io"
	"mime/multipart"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

type CustomError struct {
	Err string
}

func (e *CustomError) Error() string {
	return e.Err
}

//NewController is creating anew instance of Controlller
func NewCustomError(str string) error {
	return &CustomError{
		Err: str,
	}
}

type IntPair struct {
	first  int
	second int
}

type ICommonService interface {
	ParseStringToMatrix(str string, m int) (*mat.Dense, error)
	ReadTxtFile(textFileHeader *multipart.FileHeader) (string, error)
	ChunkSlice(slice []rune, chunkSize int, ASCII_OFFSET int) [][]float64
	ModInverse(a, m int) int
	GCD(a, b int) int
	FilterRunesAZ(runes []rune) []rune
	FilterDuplicateValues(runes []rune) []rune
	RemoveRune(runes []rune, toBeRemoved rune) []rune
	ReplaceRune(runes []rune, toBeReplaced rune, replacemenet rune) []rune
	ModLikePython(d, m int) int
}

type CommonService struct {
}

//NewCommonService is creating a new instance of CommonService
func NewCommonService() ICommonService {
	return &CommonService{}
}

func (src *CommonService) ParseStringToMatrix(str string, m int) (*mat.Dense, error) {
	data := make([]float64, m*m)

	var lines = strings.Split(str, "\n")
	if len(lines) != m {
		return nil, NewCustomError("Matrix size is not M x M")
	}

	var i int = 0
	for _, line := range lines {
		line = strings.Trim(line, "\r")
		var columns = strings.Split(line, " ")
		if len(columns) != m {
			return nil, NewCustomError("Matrix size is not M x M")
		}

		for _, column := range columns {
			floatVar, err := strconv.ParseFloat(column, 64)
			if err != nil {
				return nil, err
			}
			data[i] = floatVar
			i++
		}
	}
	ret := mat.NewDense(m, m, data)
	return ret, nil
}

func (src *CommonService) ReadTxtFile(textFileHeader *multipart.FileHeader) (string, error) {
	data := make([]byte, 256)
	file, err := textFileHeader.Open()
	if err != nil {
		return "", err
	}

	str := ""
	for {
		_, err = file.Read(data)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return "", err
			}
		}
		str += string(data)
	}

	return str, nil
}

func (src *CommonService) ChunkSlice(slice []rune, chunkSize int, ASCII_OFFSET int) [][]float64 {
	var chunks [][]float64
	var floatSlice []float64
	for _, r := range slice {
		floatSlice = append(floatSlice, float64(int(r)-ASCII_OFFSET))
	}

	for {
		if len(floatSlice) == 0 {
			break
		}

		// necessary check to avoid slicing beyond
		// slice capacity
		if len(floatSlice) < chunkSize {
			chunkSize = len(floatSlice)
		}

		chunks = append(chunks, floatSlice[0:chunkSize])
		floatSlice = floatSlice[chunkSize:]
	}

	return chunks
}

func (src *CommonService) ModInverse(a, m int) int {
	var gcdAM int
	if a > m {
		gcdAM = src.GCD(a, m)

		if gcdAM == m {
			return -1
		}
	} else {
		gcdAM = src.GCD(m, a)

		if gcdAM == a {
			return -1
		}
	}
	if gcdAM == 1 {
		return src.modInverseExtendedEuclidean(a, m)
	} else {
		return src.modInverseNaive(a, m)
	}
}

func (src *CommonService) modInverseExtendedEuclidean(A, M int) int {
	m0 := M
	y := 0
	x := 1

	if M == 1 {
		return 0
	}

	for A > 1 {
		// q is quotient
		q := A / M
		t := M

		// m is remainder now, process same as
		// Euclid's algo
		oldM := M
		M = A % M

		if M < 0 {
			M += oldM
		}

		A = t
		t = y

		// Update y and x
		y = x - q*y
		x = t
	}

	// Make x positive
	if x < 0 {
		x += m0
	}

	return x
}

func (src *CommonService) modInverseNaive(A, M int) int {
	for X := 1; X < M; X++ {
		AmM := A % M
		if AmM < 0 {
			AmM += M
		}

		XmM := X % M
		if XmM < 0 {
			XmM += M
		}

		AXM := (AmM * XmM) % M
		if AXM < 0 {
			AXM += M
		}

		if AXM == 1 {
			return X
		}
	}
	return -1
}

func (src *CommonService) GCD(a, b int) int {
	if b == 0 {
		return a
	}
	amb := a % b
	if amb < 0 {
		amb += b
	}
	return src.GCD(b, amb)
}

func (src *CommonService) FilterRunesAZ(runes []rune) []rune {
	// Ignore non-alphabetic character
	ret := []rune{}
	for i := range runes {
		if runes[i] >= 65 && runes[i] <= 90 {
			ret = append(ret, runes[i])
		}
	}
	return ret
}

func (src *CommonService) FilterDuplicateValues(runes []rune) []rune {
	keys := make(map[rune]bool)
	ret := []rune{}
	for _, entry := range runes {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			ret = append(ret, entry)
		}
	}
	return ret
}

func (src *CommonService) RemoveRune(runes []rune, toBeRemoved rune) []rune {
	ret := []rune{}
	for _, entry := range runes {
		if entry != toBeRemoved {
			ret = append(ret, entry)
		}
	}
	return ret
}

func (src *CommonService) ReplaceRune(runes []rune, toBeReplaced rune, replacemenet rune) []rune {
	for i, entry := range runes {
		if entry == toBeReplaced {
			runes[i] = replacemenet
		}
	}
	return runes
}

func (src *CommonService) ModLikePython(d, m int) int {
    var res int = d % m
    if ((res < 0 && m > 0) || (res > 0 && m < 0)) {
       return res + m
    }
    return res
}