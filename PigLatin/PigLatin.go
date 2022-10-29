package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var res string
	if len(os.Args) != 2 {
		return
	}
	arr := os.Args[1]
	for i, v := range arr {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' || v == 'A' || v == 'E' || v == 'I' || v == 'O' || v == 'U' {
			res = arr[i:] + arr[:i] + "ay"
			break
		}
	}
	if res == "" {
		res = "No vowels"
	}
	for _, v := range res {
		z01.PrintRune(v)
	}
	z01.PrintRune(10)
}
