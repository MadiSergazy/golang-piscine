package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	for _, v := range arg[0] {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
