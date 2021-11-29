package main

import "testing"

func TestSFilter(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{}},
		{[]int{5}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{3, 8}, []int{3, 8}},
		{[]int{5, 10, 7}, []int{7}},
		{[]int{10, 45, 6, 19}, []int{6, 19}},
		{[]int{21, 31, 513, 13, 35, 5}, []int{21, 31, 513, 13}},
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
