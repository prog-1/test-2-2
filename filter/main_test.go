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

func TestFilter(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{-1, 0, -3, 9, 4}, []int{-1, 4}},
		{[]int{-2, 3, 9, 21, 2}, []int{-2, 2}},
		{[]int{3, 2}, []int{2}},
	} {
		got := Filter(tc.s)
		if !equal(got, tc.want) {
			t.Errorf("Filter(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}
