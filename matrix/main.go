package main

import "fmt"

func Mutrix() (s [x][y]int) {
	for i := range s {
		for j := range s[i] {
			s[i][j] = i * j
		}
	}
	// s[0][1] = 0*1
	return s

}

func main() {
	table := Matrix()
	for i := range table {
		for j := range table[i] {
			fmt.Printf("%3d", table[i][j])
		}
		fmt.Println()
	}
}
