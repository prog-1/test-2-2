package main

import "testing"

func TestNthTerm(t *testing.T) {
	for _, tc := range []struct {
		n    uint
		want int
	}{
		{0, 1},  // base case 1
		{1, 2},  // base case 2
		{4, 38}, // from the example
	} {
		got := nthTerm(tc.n)
		if got != tc.want {
			t.Errorf("nthTerm(%v) = %v, want = %v", tc.n, got, tc.want)
		}
	}
}
