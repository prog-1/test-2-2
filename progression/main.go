package main

import "fmt"

func nthTerm(x uint) int {
	aa := 0
	ab := 2
	var y uint
	for ; y != x; y++ {
		ab++
	}
	return aa
		


	}



}
func main() {
	fmt.Println("Enter your number:")
	var x uint
	fmt.Scan(&x)
	fmt.Println(nthTerm(x))
}