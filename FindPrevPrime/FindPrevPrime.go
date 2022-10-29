package main

import "fmt"

func FindPrevPrime(nb int) int {
	if nb <= 1 {
		return 0
	}
	for a := nb; ; a-- {
		if Isprime(a) {
			return a
		}
	}
}

func Isprime(nb int) bool {
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(FindPrevPrime(2))
	fmt.Println(FindPrevPrime(4))
}
