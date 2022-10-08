package main

import (
	"bufio"
	"fmt"
	"os"
)

//ROT13 program
//rotates A-Za-z, leaves other characters untouched

func rotate(i rune) rune {
	if i >= 65 && i <= 90 {
		i += 13
		if i > 90 {
			i -= 26
		}
	}
	if i >= 97 && i <= 122 {
		i += 13
		if i > 122 {
			i -= 26
		}
	}
	return i
}

func main() {
	var rotrunes []rune
	fmt.Println("Enter text to rotate.")
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
	clearrunes := []rune(scanner.Text())
	for _, v := range clearrunes {
		u := rotate(v)
		rotrunes = append(rotrunes, u)
	}
	rottext := string(rotrunes)
	fmt.Println(rottext)
}