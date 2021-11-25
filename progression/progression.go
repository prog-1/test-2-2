package main

import "fmt"

func nthTerm(n uint) int {
	a0, a1 := 1, 2
	var i uint
	for ; i < n; i++ {
		a0, a1 = a1, 4*a0+a1
	}
	return a0
}
func main() {
	var n uint
	fmt.Scanln(&n)
	fmt.Println(nthTerm(n))
}
