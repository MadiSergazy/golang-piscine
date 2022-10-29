package main

import "fmt"

func Slice(a []string, nbrs ...int) []string {
	s := nbrs[0]
	var n int = len(a)
	if len(nbrs) == 2 {
		t := nbrs[1]
		if s > t || s < 0 && t > 0 {
			return nil
		}
		if s < 0 {
			return a[n+s : n+t]
		} else {
			return a[s:t]
		}
	} else {
		if s < 0 {
			return a[n+s:]
		}
		return a[s:]
	}
}

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}
