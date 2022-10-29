package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	len := len(arg)
	if len > 9 {
		num1 := len / 10
		num2 := len % 10
		z01.PrintRune(rune(num1) + 48)
		z01.PrintRune(rune(num2) + 48)
		z01.PrintRune('\n')
	} else {
		z01.PrintRune(rune(len + '0'))
		z01.PrintRune('\n')
	}
}
