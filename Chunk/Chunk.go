package main

import "fmt"

func Chunk(slice []int, size int) {
	var slice2 []int
	if size <= 0 {
		fmt.Println()
		return
	}
	res := make([][]int, 0, len(slice)/size+1)
	for len(slice) > size {
		slice2, slice = slice[:size], slice[size:]
		res = append(res, slice2)
	}
	if len(slice) > 0 {
		res = append(res, slice[:])
	}
	fmt.Println(res)
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
}
