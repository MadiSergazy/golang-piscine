package main

import "github.com/01-edu/z01"

func main() {
	for _, ch := range "Hello World!" {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
