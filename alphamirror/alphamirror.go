package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		for _, v := range os.Args[1] {
			if v >= 'A' && v <= 'Z' {
				v = ('Z' - v + 'A')
			} else if v >= 'a' && v <= 'z' { // val >= 97 && val <= 122
				v = ('z' - v + 'a') // 122 - 98 + 97 =121
			}
			z01.PrintRune(v)
		}
	}
	z01.PrintRune('\n')
}
