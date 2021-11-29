package main

import "fmt"

func nthTerm(n uint) int {
	a := make([]int, n+1)
	if n == 0 {
		x := 1
		return x
	}
	if n == 1 {
		x := 2
		return x
	}
	var x int
	var i uint
	a[0] = 1
	a[1] = 2
	for i = 2; i <= n; i++ {
		a[i] = 4*a[i-2] + a[i-1]
	}
	x = a[n]
	return x
}

func main() {
	fmt.Print("Enter the number:")
	var n uint
	fmt.Scan(&n)
	x := nthTerm(n)
	fmt.Println(x)
}
