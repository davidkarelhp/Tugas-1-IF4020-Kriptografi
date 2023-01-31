package services

import (
	"fmt"
	"math"
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

type IService interface {
	HillCipher(textString string, matrixString string, m int, encrypt bool) (string, error)
}

type Service struct {
}

//NewController is creating anew instance of Controlller
func NewService() IService {
	return &Service{}
}

func (src *Service) parseStringToMatrix(str string, m int) (*mat.Dense, error) {
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

func (src *Service) HillCipher(textString string, matrixString string, m int, encrypt bool) (string, error) {
	ASCII_OFFSET := 65
	numOfSymbols := 26
	matrix, err := src.parseStringToMatrix(matrixString, m)

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
	invMod := src.modInverse(detMod, numOfSymbols)
	fmt.Println("invMod = ", invMod)

	if invMod == -1 {
		return "", NewCustomError("There is no modular multiplicative inverse")
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
	runes = src.filterRunesAZ(runes)

	if len(runes)%m != 0 {
		return "", NewCustomError("Input text length should be the multiple of M")
	}

	chunks := src.chunkSlice(runes, m, ASCII_OFFSET)
	for _, chunk := range chunks {
		pmat := mat.NewDense(m, 1, chunk)
		var cmat mat.Dense
		if encrypt {
			cmat.Mul(matrix, pmat)
		} else {
			cmat.Mul(&matInvScaled, pmat)
		}

		for i := 0; i < m; i++ {
			f := math.Round(cmat.At(i, 0))
			symbolASCII := (int(f) % numOfSymbols)
			if symbolASCII < 0 {
				symbolASCII = symbolASCII + numOfSymbols
			}
			ret += string(rune(symbolASCII + ASCII_OFFSET))
		}
	}
	return ret, nil
}

func (src *Service) chunkSlice(slice []rune, chunkSize int, ASCII_OFFSET int) [][]float64 {
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

func (src *Service) modInverse(a, m int) int {
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

func (src *Service) modInverseExtendedEuclidean(A, M int) int {
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

func (src *Service) modInverseNaive(A, M int) int {
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

func (src *Service) GCD(a, b int) int {
	if b == 0 {
		return a
	}
	amb := a % b
	if amb < 0 {
		amb += b
	}
	return src.GCD(b, amb)
}

func (src *Service) filterRunesAZ(runes []rune) []rune {
	// Ignore non-alphabetic character
	ret := []rune{}
	for i := range runes {
		if runes[i] >= 65 && runes[i] <= 90 {
			ret = append(ret, runes[i])
		}
	}
	return ret
}
