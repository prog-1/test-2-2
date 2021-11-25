package main

import "fmt"

func gen(width, height int) [][]int {
	s := make([][]int, height)
	for i := range s {
		s[i] = make([]int, width)
	}

	return s
}

func main() {
	fmt.Print("Enter width and height: ")
	var width, height int
	fmt.Scan(&height, &width)
	for _, v := range gen(height, width) {
		for _, x := range v {
			fmt.Printf("%3d", x)
		}
		fmt.Println()
	}
}
