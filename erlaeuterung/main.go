package main

import (
	"fmt"
	"math/rand"
	"time"
)

// QUESTION 1: What does this function do?
// ANSWER 1: The function sorts an input slice `s` in-place in ascending order.
//
// QUESTION 2: Which algorithm does this function implement?
// ANSWER 2: Bubble sort
//
// TASK 1: Provide a better name.
func bubbleSort(s []int) {
	for i := len(s) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

// QUESTION 3: What does this function do?
// ANSWER 3: The function creates a slice that contains elements from input
// slices `a` and `b`, while maintaining the sorted order.
//
// TASK 2: Provide a better name.
func merge(a, b []int) []int {
	c := make([]int, len(a)+len(b))
	ic, ia, ib := 0, 0, 0
	for ; ia < len(a) && ib < len(b); ic++ {
		if a[ia] < b[ib] {
			c[ic] = a[ia]
			ia++
		} else {
			c[ic] = b[ib]
			ib++
		}
	}
	// QUESTION 4: Is it possible that the conditions of the following two if
	//             statements are *both* true?
	// ANSWER 4: no (the above for-loop terminates if *either* `ia < len(a)`
	// *or* `ib < len(b))`.
	//
	// QUESTION 5: Is it possible that the conditions of the following two if
	//             statements are *both* false?
	// ANSWER 5: yes (if slices `a` and `b` have the same length)
	if ia < len(a) {
		copy(c[ic:], a[ia:])
	}
	if ib < len(b) {
		copy(c[ic:], b[ib:])
	}
	return c
}

func main() {
	const maxLen, maxVal = 7, 10
	rand.Seed(time.Now().UnixNano())         // Init pseudo RNG.
	a := randInts(rand.Intn(maxLen), maxVal) // And create two vectors of random
	b := randInts(rand.Intn(maxLen), maxVal) // lengths and values.

	bubbleSort(a)
	bubbleSort(b)

	// QUESTION 6: Does the following call modify `a` and/or `b`?
	// ANSWER 6: no
	c := merge(a, b)

	fmt.Println(c)
}

// randInts return a slice of the given length filled with random numbers in the
// range [0, max).
//
// Note: this function does not need any explanation.
func randInts(l, max int) []int {
	s := make([]int, l)
	for i := range s {
		s[i] = rand.Intn(max)
	}
	return s
}
