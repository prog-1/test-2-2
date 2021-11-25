package main

import "fmt"

func filter(s []int) []int {
	var filtered []int
	for _, v := range s {
		if v%5 != 0 {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func main() {
	fmt.Print("Enter the number of elements in a slice: ")
	var size int
	fmt.Scan(&size)
	s := make([]int, size)
	fmt.Println("Enter the elements: ")
	for i := range s {
		fmt.Scan(&s[i])
	}
	fmt.Println("Filtered slice:", filter(s))
}
