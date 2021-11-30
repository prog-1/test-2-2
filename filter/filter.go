package main

import "fmt"

func filter(s []int) (d []int) {
	for _, v := range s {
		if v%5 != 0 {
			d = append(d, v)
		}
	}
	return d
}

func main() {
	fmt.Println("Enter slice sixe:")
	var size int
	fmt.Scanln(&size)
	input := make([]int, size)
	var i int
	fmt.Scanln(&i)
	for i := range input {
		fmt.Scanln(&input[i])
	}
	fmt.Println(input)
}
