package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	a := atoi(args[0])
	b := atoi(args[2])
	if a > 0 && b > 0 && (a+b) < 0 || (a < 0 && b < 0 && (a+b) > 0) {
		return
	}
	switch args[1] {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Printf("No division by 0\n")
			return
		}
		fmt.Println(a / b)
	case "%":
		if b == 0 {
			fmt.Printf("No module by 0\n")
			return
		}
		fmt.Println(a % b)
	}
}

func atoi(str string) int {
	negative := 1
	if str[0] == '-' {
		negative = -1
		str = str[1:]
	}
	num := 0
	for _, w := range str {
		if w < '0' || w > '9' {
			os.Exit(0)
		}
		num = num*10 + int(w-'0')
	}
	return negative * num
}
