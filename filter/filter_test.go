package main

import "testing"

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return true
	}
	return false
}

func TestFilter(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{}},
		{[]int{1, 4}, []int{1, 4}},
		{[]int{0}, []int{0}},
		{[]int{-6, -1}, []int{-6, -1}},
		{[]int{4, 8}, []int{}},
		{[]int{36, 23}, []int{36}},
	} {
		got := filter(tc.s)
		if !equal(got, tc.want) {
			t.Errorf("%v = (%v), want =%v", tc.s, got, tc.want)
		}
	}
}
