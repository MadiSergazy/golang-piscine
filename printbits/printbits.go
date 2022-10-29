package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		slice := make([]int, 8)
		num, err := strconv.Atoi(os.Args[1])
		if err != nil {
			for _, val := range slice {
				z01.PrintRune(rune(val) + 48)
			}
			z01.PrintRune('\n')
			return
		}
		i := len(slice) - 1
		for num != 0 {
			slice[i] = num % 2
			num /= 2
			i--
		}
		for _, val := range slice {
			z01.PrintRune(rune(val) + 48)
		}
		z01.PrintRune('\n')
	}
}
