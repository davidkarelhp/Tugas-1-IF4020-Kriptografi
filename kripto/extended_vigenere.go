package main

import (
    "bufio"
	"fmt"
    "os"
)

func modLikePython(d, m int) int {
    var res int = d % m
    if ((res < 0 && m > 0) || (res > 0 && m < 0)) {
       return res + m
    }
    return res
 }

func encryptChar(x, y int) string {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println((x + y) % 256)
	return string((x + y) % 256)
}

func decryptChar(x, y int) string {
	fmt.Println(x)
	fmt.Println(y)
    if x < y {
        modulo := modLikePython((x-y), 256)
		fmt.Println(modulo)
        return string(modulo)
    } else{
		fmt.Println((x - y) % 256)
		return string((x - y) % 256)
    }
}

func encrypt() string {
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
		char := encryptChar(int(plain[i]), int(key[j]))
		cipher = cipher + char
		j++
		if j == len(key) {
			j = 0
		}
	}
	return cipher
}

func decrypt() string {
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
		char := decryptChar(int(cipher[i]), int(key[j]))
		plain = plain + char
		j++
		if j == len(key) {
			j = 0
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