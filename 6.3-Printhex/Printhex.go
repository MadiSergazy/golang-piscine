package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*
		if len(os.Args) != 2 {
			return
		}
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			for _, v := range "ERROR" {
				z01.PrintRune(v)
			}
			return
		}
		res := strconv.FormatInt(int64(n), 16)
		for _, v := range res {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	*/

	if len(os.Args) != 2 {
		fmt.Println("ERROR")
	}
	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	res := strconv.FormatInt(int64(number), 16)
	fmt.Println(res)
}

/*
func main() {
	// arg := os.Args[1:]
	// if len(arg) != 2 {
	// 	fmt.Println("ERROR")
	// 	return
	// }
	// num, err := strconv.Atoi(arg[0])
	// if err != nil {
	// 	fmt.Println("ERROR")
	// 	return
	// }
	// hex := "0123456789abcdef"
	ConvertBase("125", "0123456789", "01")
}

func ConvertBase(num, TYPE, convTYPE string) {
	HEX := "0123456789abcdef"
	BIN := "01"
	DEL := "0123456789"
	OCT := "01234567"
	CHOUMI := "choumi"

	if convTYPE == "01" {
		num := convertDEL(num, TYPE, len(TYPE))
		convert(num, BIN, len(convTYPE))
	}
	if convTYPE == "0123456789" {
		num := convertDEL(num, TYPE, len(TYPE))
		convert(num, DEL, len(convTYPE))
	}
	if convTYPE == "01234567" {
		num := convertDEL(num, TYPE, len(TYPE))
		convert(num, OCT, len(convTYPE))
	}
	if convTYPE == "0123456789abcdef" {
		num := convertDEL(num, TYPE, len(TYPE))
		convert(num, HEX, len(convTYPE))
	}
	if convTYPE == "choumi" {
		num := convertDEL(num, TYPE, len(TYPE))
		convert(num, CHOUMI, len(convTYPE))
	}
	fmt.Println("ERROR")
}

func convertDEL(num string, TYPE string, length int) int {
	var res int
	for i, value := range num {
		for j, val := range TYPE {
			if val == value {
				res += j * int(math.Pow(float64(length), float64(len(num)-1-i)))
			}
		}
	}
	return res
}

func convert(num int, TYPE string, length int) {
	var array []int
	for num != 0 {
		array = append(array, num%length)
		num /= length
	}
	var arrayRune []rune
	for _, value := range array {
		arrayRune = append(arrayRune, rune(TYPE[value]))
	}
	for i := len(arrayRune) - 1; i > -1; i-- {
		z01.PrintRune(arrayRune[i])
	}
	z01.PrintRune('\n')
	os.Exit(0)
}
*/
