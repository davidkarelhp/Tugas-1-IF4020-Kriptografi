package main

import (
    "bufio"
	"fmt"
	"unicode"
    "os"
)

func modLikePythonAuto(d, m int) int {
    var res int = d % m
    if ((res < 0 && m > 0) || (res > 0 && m < 0)) {
       return res + m
    }
    return res
 }

func toAbjadAuto(x byte) byte {
	x = byte(unicode.ToUpper(rune(x)))
	x = x - 65
	return x
}

func encryptCharAuto(x, y byte) string {
	x = toAbjadAuto(x)
	y = toAbjadAuto(y)
	return string(((x + y) % 26) + 65)
}

func decryptCharAuto(x, y byte) string {
	x = toAbjadAuto(x)
	y = toAbjadAuto(y)
    if x < y {
        hasil := 255 - (x-y) + 1
        modulo := modLikePythonAuto(int(hasil)*(-1), 26)
        return string(modulo + 65)
    } else{
        return string(((x - y) % 26) + 65)
    }
}

func encryptAuto() string {
	var plain string
	var key string
	var cipher string = ""
	var j int = 0

	fmt.Print("Input plain text:")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    plain = scanner.Text()

	fmt.Println("Input key:")
	fmt.Scan(&key)

	for i := 0; i < len(plain); i++ {
        if plain[i] != ' '{
            char := encryptCharAuto(plain[i], key[j])
            cipher = cipher + char
            j++
            if j == len(key) {
                key = key + plain
            }
        }
	}
	return cipher
}

func decryptAuto() string {
	var cipher string
	var key string
	var plain string = ""
	var j int = 0

	fmt.Println("Input cipher text:")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    cipher = scanner.Text()

	fmt.Println("Input key:")
	fmt.Scan(&key)

	for i := 0; i < len(cipher); i++ {
        if cipher[i] != ' '{
            char := decryptCharAuto(cipher[i], key[j])
            plain = plain + char
            j++
            if j == len(key) {
                key = key + plain
            }
        }
	}
	return plain
}

func main() {
	// cipher := encryptAuto()
	// fmt.Println(cipher)
	plain := decryptAuto()
	fmt.Println(plain)
}