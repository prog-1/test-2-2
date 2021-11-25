package main

import "fmt"

func filter(s []int) []int {
	var a []int
	for _, v := range s {
		if v%5 != 0 {
			a = append(a, v)
		}
	}
	return a
}

func main() {
	fmt.Println("Enter slice size:")
	var size uint
	fmt.Scan(&size)
	input := make([]int, size)
	for i := range input {
		fmt.Scan(&input[i])
	}
	a := filter(input)
	fmt.Println("The filtered slice:", a)
}
