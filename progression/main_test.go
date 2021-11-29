package main

import "testing"

func TestAnyIsEven(t *testing.T) {
	for _, tc := range []struct {
		n    uint
		want int
	}{
		{0, 1},
		{1, 2},
		{2, 6},
		{3, 14},
		{4, 38},
	} {
		got := nthTerm(tc.n)
		if got != tc.want {
			t.Errorf("nthTerm(%v) = %v, want = %v", tc.n, got, tc.want)
		}
	}
}
