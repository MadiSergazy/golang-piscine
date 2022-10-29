package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 || Atoi(os.Args[1]) == 0 || string(os.Args[1][0]) == "-" {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	n := (Atoi(os.Args[1]))
	sum := 0
	for i := n; i >= 2; i-- {
		if Isprime(i) {
			sum += i
		}
	}
	for _, v := range Itoa(sum) {
		z01.PrintRune(v)
	}
	z01.PrintRune(10)
}

func Isprime(n int) bool {
	for a := 2; a <= n/2; a++ {
		if n%a == 0 {
			return false
		}
	}
	return true
}

func Atoi(s string) int {
	n := 0
	for _, v := range s {
		n = n*10 + int(v-'0')
	}
	return n
}

func Itoa(n int) string {
	digits := "0123456789"
	if n < 10 {
		return digits[n : n+1]
	}
	return Itoa(n/10) + digits[n%10:n%10+1]
}
