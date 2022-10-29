package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var res []string
	word := ""
	for _, v := range os.Args[1] {
		if v != ' ' {
			word += string(v)
		} else if word != "" {
			res = append(res, word)
			word = ""
		} else {
			continue
		}
	}
	if word != "" {
		res = append(res, word)
	}
	val := ""
	for i := len(res) - 1; i > 0; i-- {
		val = res[i] + " " + val
	}
	val += res[0]
	for _, v := range val {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

/*
func main() {
	if len(os.Args) == 2 {
		str := Split(os.Args[1])

		for _, val := range str[1:] {
			Print(val)
			z01.PrintRune(' ')
		}
		Print(str[0])
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('\n')
	}
}

func Print(s string) {
	for _, val := range s {
		z01.PrintRune(val)
	}
}

func Split(str string) []string {
	var res []string
	word := ""

	for _, val := range str {
		if val == ' ' && word != "" {
			res = append(res, word)
			word = ""
		} else if val != ' ' {
			word += string(val)
		}
	}
	if word != "" {
		res = append(res, word)
	}
	return res
}
*/
