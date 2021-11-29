package main

import "fmt"

func nthTerm(n uint) int {
	a, b := 0, 2
	var tmp uint
	for ; tmp != n; tmp++ {
		b++
	}
	return a
}

func main() {
	fmt.Print("Enter n: ")
	var n uint
	fmt.Scan(&n)
	fmt.Println(nthTerm(n))
}
