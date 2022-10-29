package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 4 {
		s1 := os.Args[2]
		s2 := os.Args[3]
		for _, v := range os.Args[1] {
			if v == rune(s1[0]) {
				z01.PrintRune(rune(s2[0]))
			} else {
				z01.PrintRune(v)
			}
		}
		z01.PrintRune(10)
	}
}
