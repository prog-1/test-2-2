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
		{[]int{0}, []int{},
		{[]int{5}, []int{}},
		{[]int{8, 7, 10}, []int{8, 7}},
		{[]int{5, 10, 13}, []int{13}},
	} {
		got := Filter(tc.s)
		if !equal(got, tc.want) {
			t.Errorf("Filter(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}
