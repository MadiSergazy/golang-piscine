package main

import (
	"strconv"

	"github.com/01-edu/z01"
)

func ReduceInt(a []int, f func(int, int) int) {
	n := a[0]
	for i := 1; i < len(a); i++ {
		n = f(n, a[i])
	}
	arr := strconv.Itoa(n)
	for _, v := range arr {
		z01.PrintRune(v)
	}
	z01.PrintRune(10)
}

// func Itoa(n int) string {
// 	num := n
// 	res := ""
// 	if n < 0 {
// 		num = -n
// 		res = "-"
// 	}
// 	digits := "0123456789"
// 	if num < 10 {
// 		return res + digits[num:num+1]
// 	}
// 	return res + Itoa(num/10) + digits[num%10:num%10+1] // res- +Itoa(n/10) +digits[144%10 : 144%10 +1]==4
// }
