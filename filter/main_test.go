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

func TestFilterOdds(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{}},
		{[]int{2, 3, 7, 10}, []int{2, 3, 7}},
		{[]int{3, 5, 8}, []int{3, 8}},
		{[]int{-5, 5}, []int{}},
		{[]int{-10, -7, 4}, []int{-7, 4}},
	} {
		got := make([]int, len(tc.s))
		copy(got, tc.s)
		got = filter(got)
		if !equal(got, tc.want) {
			t.Errorf("Filter(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}
