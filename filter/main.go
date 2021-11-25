package main

import "fmt"

func Filter(s []int) []int {
	x := make([]int, 0)
	for _, v := range s {
		if v%3 != 0 {
			x = append(x, v)
		}
	}
	return x
}

func main() {
	fmt.Println("Enter slice size:")
	var size uint
	fmt.Scan(&size)
	input := make([]int, size)
	for i := range input {
		fmt.Scan(&input[i])
	}
	fmt.Println("FilterEvens", Filter(input))
}
