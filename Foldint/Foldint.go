package main

import (
	"github.com/01-edu/z01"
)

func FoldInt(f func(int, int) int, a []int, n int) {
	for _, num := range a {
		n = f(n, num)
	}
	arr := Itoa(n)
	for _, v := range arr {
		z01.PrintRune(v)
	}
	z01.PrintRune(10)
}

func Itoa(n int) string { // 144 - 14
	num := n
	ch := ""
	if n == 0 {
		return "0"
	}
	if n < 0 {
		num = -n
		ch = "-"
	}
	digits := "0123456789"
	if num < 10 {
		return ch + digits[num:num+1]
	}
	return ch + Itoa(num/10) + digits[num%10:num%10+1] // ch - + Itoa(n/10) + digits[144%10 : 144%10 + 1]== 4
}

func Add(n, num int) int {
	return n + num
}

func Mul(n, num int) int {
	return n * num
}

func Sub(n, num int) int {
	return n - num
}

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}
