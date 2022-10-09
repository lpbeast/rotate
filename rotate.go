package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	in := os.Stdin
	out := os.Stdout
	var err error
	argcount := len(os.Args[1:])
	if argcount > 0 {
		for i := 1; i <= argcount; i++ {
			parsed := strings.Split(os.Args[i], "=")
			//if multiple input or output files are specified, only the first one given will be used
			switch {
			case parsed[0] == "-i":
				if in == os.Stdin {
					in, err = os.Open(parsed[1])
					if err != nil {
						fmt.Println(err)
						fmt.Println("Exiting.")
						return
					}
					defer in.Close()
				}
			case parsed[0] == "-o":
				if out == os.Stdout {
					out, err = os.Create(parsed[1])
					if err != nil {
						fmt.Println(err)
						fmt.Println("Exiting.")
						return
					}
					defer out.Close()
				}
			default:
				fmt.Println("Invalid option")
				fmt.Println("Valid options are -i=filename for an input file and -o=filename for an output file.")
				return
			}
		}
	}
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		fmt.Fprintln(out, strprocess(scanner.Text()))
	}
}