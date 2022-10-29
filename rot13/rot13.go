package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		for _, ch := range os.Args[1] {
			if ch >= 'A' && ch <= 'Z' {
				if ch >= 'A'+13 {
					ch -= 13
				} else {
					ch += 13
				}
			}
			if ch >= 'a' && ch <= 'z' {
				if ch >= 'a'+13 {
					ch -= 13
				} else {
					ch += 13
				}
			}
			z01.PrintRune(ch)
		}
		z01.PrintRune(10)
	}
}
