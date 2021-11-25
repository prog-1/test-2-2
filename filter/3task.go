package main

import "fmt"

func filter(s []int) []int {
	result := make([]int, 0)
	for _, v := range s {
		if v%5 != 0 {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	fmt.Println("Enter slice size:")
	size := 0
	fmt.Scan(&size)
	fmt.Println("Enter the elements:")
	input := make([]int, size)
	for i := range input {
		fmt.Scan(&input[i])
	}
	fmt.Println("The filtered slice:", filter(input))