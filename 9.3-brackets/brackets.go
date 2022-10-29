package main

import (
	"fmt"
	"os"
)

func main() {
	ar := os.Args
	if len(ar) > 1 {
		for _, x := range ar[1:] {
			fmt.Println(Brackets1(x))
		}
	} else {
		fmt.Println()
	}
}
func Brackets1(s string) string {
	//mapa := make(map[string]bool)
	br := ""
	for _, v := range s{
		if v == ']' || v == '}' || v == ')' && len(br) == 0{
			return "ERROR"
		}
		if v == '[' || v == '{' || v == '('{
			br += string(v)
		}
		if v == ']' && br[len(br)-1] == '['{
			br = br[:len(br)-1]
		}
		if v == ')' && br[len(br)-1] == '('{
			br = br[:len(br)-1]
		}
		if v == '}' && br[len(br)-1] == '{'{
			br = br[:len(br)-1]
		}

	}
	if len(br) == 0{
		return "OK"
	} else {
		return "ERROE"
	}


}

func Brackets(s string) string {
	br := ""
	for _, r := range s { // базовый случай
		if (r == ')' || r == '}' || r == ']') && len(br) == 0 {
			return "Error"
		}
		if r == '(' || r == '{' || r == '[' { //случи первый при находке входа
			br += string(r)
		}
		if r == ')' && br[len(br)-1] == '(' { //при закрытиии скобок
			br = br[:len(br)-1]
		}
		if r == '}' && br[len(br)-1] == '{' { // при закрытиии скобок
			br = br[:len(br)-1]
		}
		if r == ']' && br[len(br)-1] == '[' {
			br = br[:len(br)-1]
		}
	}
	if len(br) == 0 { // если в конце у нас неостанется элементов тогла все ок
		return "OK"
	}
	return "Error"
}
