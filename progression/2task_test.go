package main

import "testing"

func TestNthTerm(t *testing.T) {
	for _, tc := range []struct {
		x    uint
		want int
	}{
		{0, 1},
		{4, 38},
		{1, 2},
	} {
		got := nthTerm(tc.x)
		if got != tc.want {
			t.Errorf("nthTerm(%v) = %v, want = %v", tc.x, got, tc.want)
		}
	}
}
