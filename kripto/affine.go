package main

import (
    "bufio"
	"fmt"
	"unicode"
    "os"
)

func modInverse(x int) int{
	for i := 1; i < 26; i++{
		if (((x % 26) * (i % 26)) % 26 == 1){
			return i
		}
	}
	return -1
}

func contains(x int) bool {
	var arr = [12]int{1, 3, 5, 7, 9, 11, 15, 17, 19, 21, 23, 25}
	for _, v := range arr {
		if v == x {
			return true
		}
	}

	return false
}

func checkAbjad(x byte) bool{
    if (x >=65 && x <=90) || (x >= 97 && x <=122){
        return true
    } else {
        return false
    }
}

func modLikePython(d, m int) int {
    var res int = d % m
    if ((res < 0 && m > 0) || (res > 0 && m < 0)) {
       return res + m
    }
    return res
 }

func toAbjad(x byte) byte {
	x = byte(unicode.ToUpper(rune(x)))
	x = x - 65
	return x
}

func encryptChar(x byte, m,b int) string {
	x = toAbjad(x)
	return string(((m*int(x) + b) % 26) + 65)
}

func decryptChar(x byte, m,b int) string {
	x = toAbjad(x)
	m = modInverse(m)
	hasil := m*(int(x)-b)
	fmt.Println(-120 % 26)
	if hasil < 0{
		return string(modLikePython(hasil, 26) + 65)
	}
    return string(((m*(int(x)-b)) % 26) + 65)
}

func encrypt() string {
	var plain string
	var keyM int
	var keyB int
	var cipher string = ""

	fmt.Print("Input plain text:")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    plain = scanner.Text()

	for{
		fmt.Println("Input key m (relatif prima dengan 26):")
		fmt.Scan(&keyM)

		if contains(keyM){
			break
		}
		fmt.Println("Not valid input! Try again.")
	}

	fmt.Println("Input key b (pergeseran):")
	fmt.Scan(&keyB)

	for i := 0; i < len(plain); i++ {
        if checkAbjad(plain[i]){
            char := encryptChar(plain[i], keyM, keyB)
            cipher = cipher + char
        }
	}
	return cipher
}

func decrypt() string {
	var cipher string
	var keyM int
	var keyB int
	var plain string = ""

	fmt.Println("Input cipher text:")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    cipher = scanner.Text()

	for{
		fmt.Println("Input key m (relatif prima dengan 26):")
		fmt.Scan(&keyM)

		if contains(keyM){
			break
		}
		fmt.Println("Not valid input! Try again.")
	}

	fmt.Println("Input key b (pergeseran):")
	fmt.Scan(&keyB)

	for i := 0; i < len(cipher); i++ {
        if checkAbjad(cipher[i]){
            char := decryptChar(cipher[i], keyM, keyB)
            plain = plain + char

        }
	}
	return plain
}

func main() {
	// cipher := encrypt()
	// fmt.Println(cipher)
	plain := decrypt()
	fmt.Println(plain)
}
