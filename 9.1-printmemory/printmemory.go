package main

import (
	"unicode"

	"github.com/01-edu/z01"
)

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func tohex(n int) string {
	if n == 0 {
		return "00"
	}

	base := []rune("0123456789abcdef")

	res := ""

	for n > 0 {
		a := n % 16
		res = string(base[a]) + res
		n = n / 16
	}

	if len(res) == 1 {
		return "0" + res
	}

	return res
}

func PrintMemory(arr [10]byte) {
	var a string

	for _, v := range arr {
		a += tohex(int(v)) + " "
	}

	print(a[:11])
	print(a[12:23])
	print(a[24:29])
	str := ""
	for _, v := range arr {
		if unicode.IsGraphic(rune(v)) {
			str += string(v)
			// fmt.Println("byte:", v)
		} else {
			str += "."
		}
	}

	print(str)
}

func print(str string) {
	for _, j := range str {
		z01.PrintRune(j)
	}
	z01.PrintRune('\n')
}
