package main

import "fmt"

func nthTerm(n uint) int {
	a, b := 1, 2
	for i := uint(0); i < n; i++ {
		a, b = b, 4*a+b
	}
	return a
}

func main() {
	fmt.Print("Enter n: ")
	var n uint
	fmt.Scan(&n)
	fmt.Println("The nth term is", nthTerm(n))
}
