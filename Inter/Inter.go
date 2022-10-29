package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]
	res := ""
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if v1 == v2 && Contain(res, v1) {
				res += string(v1)
			}
		}
	}
	for _, v := range res {
		z01.PrintRune(v)
	}
	z01.PrintRune(10)
}

func Contain(s string, r rune) bool {
	for _, v := range s {
		if v == r {
			return false
		}
	}
	return true
}
