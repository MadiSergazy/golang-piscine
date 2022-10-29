package main

import (
	"fmt"
	"os"
)

func main() {
	// if len(os.Args) != 2 {
	// 	return
	// }
	// if Atoi(os.Args[1]) >= 4000 || Atoi(os.Args[1]) == 0 {
	// 	fmt.Printf("ERROR: cannot convert to roman digit\n")
	// 	return
	// }
	first, second := Roman(Atoi("1655"))
	fmt.Printf("%s\n", first)
	fmt.Printf("%s\n", second)
}

func Roman(n int) (string, string) {
	array_num := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	array_roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	array_output := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}
	var res, roman string
	i := 0
	for n > 0 {
		k := n / array_num[i]
		for j := 0; j < k; j++ {
			res += array_output[i] + "+"
			roman += array_roman[i]
			n -= array_num[i]
		}
		i++
	}
	return res[:len(res)-1], roman
}

func Atoi(s string) int {
	n := 0
	for _, v := range s {
		if v < '0' || v > '9' {
			fmt.Printf("ERROR: cannot convert to roman digit\n")
			os.Exit(0)
		}
		n = n*10 + int(v-'0')
	}
	return n
}
