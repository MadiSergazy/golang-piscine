package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	res := ""
	if len(os.Args) == 3 {
		str1 := os.Args[1]
		str2 := os.Args[2]
		for _, v := range str1 {
			if Contain(res, v) {
				res += string(v)
			}
		}
		for _, v := range str2 {
			if Contain(res, v) {
				res += string(v)
			}
		}
		for _, v := range res {
			z01.PrintRune(v)
		}
		z01.PrintRune(10)
	} else {
		z01.PrintRune(10)
	}
}

func Contain(s string, r rune) bool {
	for _, v := range s {
		if v == r {
			return false
		}
	}
	return true
}
