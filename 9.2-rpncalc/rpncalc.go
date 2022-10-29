package main

import (
	"fmt"
	"strconv"
)

func main() {
	//if len(os.Args) != 2 {
	//		return
	//	}
	//s := os.Args[1]
	s := "1 2 * 3 * 4 +"

	var newStr []string
	for _, v := range s {
		if v != ' ' {
			newStr = append(newStr, string(v))
		}
	}
	fmt.Println(newStr)

	var arr []int
	for _, v := range newStr {
		num, err := strconv.Atoi(v)
		if err != nil {
			if len(arr) < 2 {
				fmt.Println("ERROR1")
				return
			}
			switch v {
			case "*":
				arr[len(arr)-2] = arr[len(arr)-2] * arr[len(arr)-1]
				arr = arr[:len(arr)-1]
			case "/":
				arr[len(arr)-2] = arr[len(arr)-2] / arr[len(arr)-1]
				arr = arr[:len(arr)-1]
			case "-":
				arr[len(arr)-2] = arr[len(arr)-2] - arr[len(arr)-1]
				arr = arr[:len(arr)-1]
			case "+":
				arr[len(arr)-2] = arr[len(arr)-2] + arr[len(arr)-1]
				arr = arr[:len(arr)-1]
			case "%":
				arr[len(arr)-2] = arr[len(arr)-2] % arr[len(arr)-1]
				arr = arr[:len(arr)-1]
			case "":
				continue
			default:
				fmt.Println("ERROR2")
				return

			}
		} else { //есни нет оибок конвертировки тогда берем и добовим
			arr = append(arr, num)
		}
	}
	if len(arr) != 1 {
		fmt.Println(arr[0])
		fmt.Println("ERROR3")
		return
	}
	fmt.Println(arr[0])
}

/*
	input := strings.Split(s, " ")
	var arr []int
	newinput := []string{}
	for _, item := range input {
		if item != "" {
			newinput = append(newinput, item)
		}
	}
	for _, item := range newinput {
		num, err := strconv.Atoi(item)
		if err != nil {
			if len(arr) < 2 {
				fmt.Print("Error\n")
				return
			}
			switch item {
			case "+":
				arr[len(arr)-2] = arr[len(arr)-2] + arr[len(arr)-1]
				arr = arr[:len(arr)-1]
			case "-":
				arr[len(arr)-2] = arr[len(arr)-2] - arr[len(arr)-1]
				arr = arr[:len(arr)-1]
			case "/":
				arr[len(arr)-2] = arr[len(arr)-2] / arr[len(arr)-1]
				arr = arr[:len(arr)-1]
			case "%":
				arr[len(arr)-2] = arr[len(arr)-2] % arr[len(arr)-1]
				arr = arr[:len(arr)-1]
			case "*":
				arr[len(arr)-2] = arr[len(arr)-2] * arr[len(arr)-1]
				arr = arr[:len(arr)-1]
			case "":
				continue
			default:
				fmt.Println("Error")
				return
			}
		} else {
			arr = append(arr, num)
		}
	}
	if len(arr) != 1 {
		fmt.Println("Error")
		return
	}
	fmt.Println(arr[0])
}
*/
