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

func strprocess(s string) string {
	var rotrunes []rune
	clearrunes := []rune(s)
	for _, v := range clearrunes {
		u := rotate(v)
		rotrunes = append(rotrunes, u)
	}
	return string(rotrunes)
}

func main() {
	argcount := len(os.Args[1:])
	switch {
	case argcount == 0:
		//Accept a line of text from console, rotate it, display output on console
		fmt.Println("Enter text to rotate.")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		rottext := strprocess(scanner.Text())
		fmt.Println(rottext)
		
	case argcount == 1:
		//Read the given file, rotate it, display output on console
		fmt.Println("Input file is", os.Args[1])
		
		infile, err := os.Open(os.Args[1])
		defer infile.Close()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Exiting.")
			return
		}

		scanner := bufio.NewScanner(infile)
		for scanner.Scan() {
			rottext := strprocess(scanner.Text())
			fmt.Println(rottext)
		}
		
	case argcount == 2:
		//Read the first file named, rotate it, write output to the second file named.
		fmt.Println("Input file is", os.Args[1])
		fmt.Println("Output file is", os.Args[2])
		
		infile, err := os.Open(os.Args[1])
		defer infile.Close()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Exiting.")
			return
		}
		
		outfile, err := os.Create(os.Args[2])
		defer outfile.Close()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Exiting.")
			return
		}
		
		scanner := bufio.NewScanner(infile)
		for scanner.Scan() {
			rottext := strprocess(scanner.Text()) + "\n"
			outfile.WriteString(rottext)
		}
		
	default:
		fmt.Println("Too many arguments.")
	}
}