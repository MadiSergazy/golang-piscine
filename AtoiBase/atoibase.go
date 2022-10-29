package main

import (
	"fmt"
)

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	//1 умножаем первое чисо на длину того что надо перевести (начинакм с нуля чтобы незатупить )
	// добовляем его позицию
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "-0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}
func AtoiBase(s, base string) int {
	mapa, res, lenBase := make(map[rune]int), 0, len(base)
	for i, v := range base {
		if v == '+' || v == '-' {
			return 0
		}
		mapa[v] = i
	}
	for _, v := range s {

		res = res*lenBase + mapa[v]
	}
	return res
}

/*
func AtoiBase(s, base string) int {
	m := make(map[rune]int)
	for i, v := range base {
		if v == '+' || v == '-' {
			return 0
		}
		m[v] = i
	}
	var num int
	for _, v2 := range s {
		num = num*len(base) + m[v2]
	}
	return num
}*/

// func AtoiBase(s string, t string) int {
// 	res := 0
// 	lens := 0
// 	arr := map[rune]bool{}
// 	for _, c := range t {
// 		if arr[c] || c == '-' || c == '+' {
// 			return res
// 		}
// 		arr[c] = true
// 		lens++
// 	}
// 	if lens > 1 {
// 		for _, v := range s {
// 			count := 0
// 			if arr[v] {
// 				for _, j := range t {
// 					if j == v {
// 						break
// 					}
// 					count++
// 				}
// 				res = res*lens + count
// 			}
// 		}
// 	}
// 	return res
// }
