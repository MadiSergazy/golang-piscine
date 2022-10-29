package main

import (
	"fmt"
)

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"} // 1 2 3 A B C a b c
	SortWordArr(result)
	fmt.Println(result)
}

func SortWordArr(a []string) {
	/*for i := 0; i < len(a); i++ { // a
		for j := i; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}*/

	/*for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}*/
	var i, j int
	for i, j = 0, i; i < len(a); {
		if a[i] > a[j] {
			a[i], a[j] = a[j], a[i]
		}
		i++
		j++
	}
}
