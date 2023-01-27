package main

import (
    "bufio"
	"fmt"
	"unicode"
    "os"
)

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

func encryptChar(x, y byte) string {
	x = toAbjad(x)
	y = toAbjad(y)
	return string(((x + y) % 26) + 65)
}

func decryptChar(x, y byte) string {
	x = toAbjad(x)
	y = toAbjad(y)
    if x < y {
        hasil := 255 - (x-y) + 1
        modulo := modLikePython(int(hasil)*(-1), 26)
        return string(modulo + 65)
    } else{
        return string(((x - y) % 26) + 65)
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

	fmt.Print("Input key:")
	fmt.Scan(&key)

	for {
		if len(key) <= len(plain) {
			break
		} else {
			fmt.Println("Invalid key, must be shorter than plain text")
			fmt.Println("Try again")
			fmt.Print("Input key:")
			fmt.Scan(&key)
		}
	}

	for i := 0; i < len(plain); i++ {
        if plain[i] != ' '{
            char := encryptChar(plain[i], key[j])
            cipher = cipher + char
            j++
            if j == len(key) {
                j = 0
            }
        }
	}
	return cipher
}

func decrypt() string {
	var cipher string
	var key string
	var plain string = ""
	var j int = 0

	fmt.Print("Input cipher text:")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    cipher = scanner.Text()

	fmt.Print("Input key:")
	fmt.Scan(&key)

	for {
		if len(key) <= len(cipher) {
			break
		} else {
			fmt.Println("Invalid key, must be shorter than plain text")
			fmt.Println("Try again")
			fmt.Print("Input key:")
			fmt.Scan(&key)
		}
	}

	for i := 0; i < len(cipher); i++ {
        if cipher[i] != ' '{
            char := decryptChar(cipher[i], key[j])
            plain = plain + char
            j++
            if j == len(key) {
                j = 0
            }
        }
	}
	return plain
}

func main() {
	var endec int

	fmt.Print("Input 1 for encrypt, 2  for decrypt")
	fmt.Scan(&endec)

    if endec == 1{
        cipher := encrypt()
        fmt.Println(cipher)
    } else {
        plain := decrypt()
        fmt.Println(plain)
    }
}
