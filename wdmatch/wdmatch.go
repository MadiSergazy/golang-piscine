package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Match(r rune, s string) int {
	for i, ch := range s {
		if ch == r {
			return i
		}
	}
	return -1
}

func main() {
	word := os.Args[1]
	base := os.Args[2]
	var res string
	for _, ch := range word {
		if n := Match(ch, base); n == -1 {
			return
		} else {
			res += string(ch)
			base = base[n+1:]
		}
	}
	for _, ch := range res {
		z01.PrintRune(ch)
	}
	z01.PrintRune(10)
}
