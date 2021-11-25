package main

import "testing"

func TestNthTerm(t *testing.T) {
	for _, tc := range []struct {
		n    uint
		want int
	}{
		{0, 1},
		{4, 38},
		{1, 2},
	} {
		got := nthTerm(tc.n)
		if got != tc.want {
			t.Errorf("nthTerm(%v) = %v, want = %v", tc.n, got, tc.want)
		}
	}
}
