package main

import "fmt"

func filter(s []int) (res []int) {
	for _, v := range s {
		if v%5 != 0 {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	fmt.Print("Enter the number of elements in a slice: ")
	var size int
	fmt.Scan(&size)
	fmt.Print("Enter the elements: ")
	s := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&s[i])
	}
	fmt.Println("The filtered slice:", filter(s))
}
