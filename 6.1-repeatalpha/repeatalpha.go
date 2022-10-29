package main

import (
	"fmt"
	"os"
	//"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		for _, v := range os.Args[1] {
			if v >= 'a' && v <= 'z' {
				for i := 0; i <= int(v-'a'); i++ {
					//z01.PrintRune(v)
					fmt.Print(string(v))
				}
			} else if v >= 'A' && v <= 'Z' {
				for i := 0; i <= int(v-'A'); i++ {
					//z01.PrintRune(v)
					fmt.Print(string(v))
				}
			} else {
				//z01.PrintRune(v)
				fmt.Print(string(v))
			}
		}
		//	z01.PrintRune('\n')
		fmt.Print("\n")
	}
}
