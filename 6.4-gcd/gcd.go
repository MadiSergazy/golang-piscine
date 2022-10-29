package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*
		if len(os.Args) == 3 {
			num1, err1 := strconv.Atoi(os.Args[1])
			num2, err2 := strconv.Atoi(os.Args[2])
			if err1 != nil && err2 != nil {
				return
			}
			res := 0
			for i := 1; i <= num1 && i <= num2; i++ {
				if num1%i == 0 && num2%i == 0 {
					res = i
				}
			}
			fmt.Println(res)
		}
	}*/

	if len(os.Args) == 3 {

		num1, err1 := strconv.Atoi(os.Args[1])
		num2, err2 := strconv.Atoi(os.Args[2])
		if err1 != nil || err2 != nil {
			return
		}
		res := 0
		for i := 1; i <= num1 && i <= num2; i++ {
			if num1%i == 0 && num2%i == 0 {
				res = i

			}
		}
		fmt.Println(res)
	}
}

/*if len(os.Args) == 3 {
	num1, err1 := strconv.Atoi(os.Args[1])
	num2, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil && err2 != nil {
		return
	}
	res := 0
	for i := 1; i <= num1 && i <= num2; i++ {
		if num1%i == 0 && num2%i == 0 {
			res = i
		}
	}
	fmt.Println(res)
}*/
