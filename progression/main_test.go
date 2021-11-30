package main

import "testing"

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestnthTerm(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{0}, []int{1}},
		{[]int{2}, []int{6}},
		{[]int{3}, []int{14}},
	} {
		got := nthTerm(tc.s)
		if !equal(got, tc.want) {
			t.Errorf("nthTerm(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}
