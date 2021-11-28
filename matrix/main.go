package main

import "fmt"

func gen(width, height int) [][]int {
	if width <= 0 || height <= 0 {
		return nil
	}

	m := make([][]int, height)
	for row := 0; row < height; row++ {
		m[row] = make([]int, width)
	}

	m[height-1][width-1] = 1
	for i := height - 2; i >= 0; i-- {
		for j := width - 1; j >= 0; j-- {
			for r := i + 1; r < height; r++ {
				m[i][j] += m[r][j]
			}
			for c := j + 1; c < width; c++ {
				m[i][j] += m[i][c]
			}
		}
	}
	return m
}

func main() {
	var width, height int
	fmt.Print("Enter width and heigth: ")
	fmt.Scan(&width, &height)
	matrix := gen(width, height)
	fmt.Println("Result:")
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Printf("%4d", matrix[i][j])
		}
		fmt.Println()
	}
}
