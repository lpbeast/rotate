package main

import (
	"bufio"
	"fmt"
	"os"
)

//ROT13 program
//rotates A-Za-z, leaves other characters untouched

func rotate(i rune) rune {
	if i >= 'A' && i <= 'Z' {
		i += 13
		if i > 'Z' {
			i -= 26
		}
	}
	if i >= 'a' && i <= 'z' {
		i += 13
		if i > 'z' {
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