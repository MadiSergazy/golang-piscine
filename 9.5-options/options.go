package main

import (
	"fmt"
	"os"
)

func main() {
	table := "00000000 00000000 00000000 00000000"
	options := "******zy xwvutsrq ponmlkji hgfedcba"
	help := "options: abcdefghijklmnopqrstuvwxyz"
	invalid := "Invalid Option"
	if len(os.Args) == 1 {
		fmt.Println(help)
		return
	}
	for _, k := range os.Args[1:] {
		if k[0] == '-' && len(k) == 1 {
			fmt.Println(invalid)
			return
		}
		k = k[1:]
		for j, l := range k {
			if l == 'h' && j == 0 {
				fmt.Println(help)
				return
			}
			for i, m := range options {
				if l >= 'a' && l <= 'z' || (l == '-' && j != len(k)-1) {
					if l == m {
						table = table[:i] + "1" + table[i+1:]
						break
					}
				} else {
					fmt.Println(invalid)
					return
				}
			}
		}
	}
	fmt.Println(table)
}
