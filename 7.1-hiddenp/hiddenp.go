package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		return
	}
	res, s1, s2, index := "", os.Args[1], os.Args[2], 0
	for _, v := range s2 {
		if len(res) == len(s1) {
			break
		}
		if v == rune(s1[index]) {
			res += string(v)
		}
	}
	if res == s1 {
		fmt.Println(res)
	} else {
		fmt.Println("0")
	}

	/*
		if len(os.Args) == 3 {



			s1 := os.Args[1]
			s2 := os.Args[2]
			i := 0
			res := ""
			for _, v := range s2 {
				if len(res) == len(s1) {
					break
				}
				if v == rune(s1[i]) {
					res += string(v)
					i++
				}
			}
			if res == s1 {
				fmt.Println(1)
				return
			}
			fmt.Println(0)
		}*/
}
