package main

import "testing"

func TestFilter(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{10, -15}, []int{}},
		{[]int{-1, 0, -5, 15, 4}, []int{-1, 4}},
		{[]int{-2, -3}, []int{-2, -3}},
		{[]int{8, 9, 10}, []int{8, 9}},
	} {
		got := filter(tc.s)
		if !equal(got, tc.want) {
			t.Errorf("filter(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}

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
