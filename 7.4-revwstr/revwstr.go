package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	var answer []string
	word := ""
	for _, v := range args[0] {
		switch v {
		case ' ':
			answer = append(answer, word)
			word = ""
		default:
			word += string(v)

		}
		/*
			if v != ' ' {
				word += string(v)
			} else if v == ' ' {
				answer = append(answer, word)
				word = ""
			}*/
	}
	if word != "" {
		answer = append(answer, word)
		word = ""
	}

	for i := len(answer) - 1; i >= 0; i-- {
		for _, l := range answer[i] {
			z01.PrintRune(l)
		}
		if i != 0 {
			z01.PrintRune(' ')
		}

	}
}

/*
{
	// if len(os.Args) == 2 {
	var res []string
	word := ""
	for _, v := range os.Args[1] {
		if v != ' ' {
			word += string(v)
		}
		if v == ' ' {
			res = append(res, word)
			word = ""
		}
	}
	if word != "" {
		res = append(res, word)
	}
	for i := len(res) - 1; i >= 0; i-- {
		for _, j := range res[i] {
			z01.PrintRune(j)
		}
		if i != 0 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
	// }
}
*/
// func main() {
// 	if len(os.Args) != 2 {
// 		return
// 	}
// 	s := os.Args[1]
// 	var res string
// 	for i := len(s) - 1; i != -1; i-- {
// 		if s[i] != ' ' {
// 			res = string(s[i]) + res
// 		}
// 		if s[i] == ' ' {
// 			if res != " " {
// 				Myprint(res + " ")
// 			}
// 			res = ""
// 		}
// 	}
// 	if len(res) != 0 {
// 		Myprint(res)
// 	}
// 	z01.PrintRune(10)
// }

// func Myprint(s string) {
// 	for _, v := range s {
// 		z01.PrintRune(v)
// 	}
// }
