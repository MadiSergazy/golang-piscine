package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		n, ea := strconv.Atoi(os.Args[1])
		if ea != nil {
			z01.PrintRune('\n')
			return
		}
		if isPowerofTwo(n) {
			for _, r := range "true" {
				z01.PrintRune(r)
			}
		} else {
			for _, r := range "false" {
				z01.PrintRune(r)
			}
		}
	}
	z01.PrintRune('\n')
}

func isPowerofTwo(n int) bool {
	if n&(n-1) == 0 && n != 0 { // битовая операция //1000 - 8 в десятичной системе
		// 0111 - 7 в десятичной системе
		return true
	}
	return false
}
