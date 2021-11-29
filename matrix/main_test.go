package main

import (
	"testing"
)

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) > 0 && len(a[0]) != len(b[0]) {
		return false
	}
	for i := range a {
		for j := range a[0] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestMatrix(t *testing.T) {
	for _, tc := range []struct {
		width, height int
		want          [][]int
	}{
		// one of dimensions is zero
		{0, 0, nil},
		{1, 0, nil},
		{0, 1, nil},
		// negative dimensions
		{-2, 0, nil},
		{0, -3, nil},
		// trivial case
		{1, 1, [][]int{
			{1},
		}},
		// case from the example
		{5, 4, [][]int{
			{130, 53, 21, 8, 4},
			{36, 16, 7, 3, 2},
			{8, 4, 2, 1, 1},
			{0, 0, 0, 0, 1},
		}},
	} {
		got := gen(tc.width, tc.height)
		if !equal(got, tc.want) {
			t.Errorf("gen(%v, %v) = %v, want = %v", tc.width, tc.height, got, tc.want)
		}
	}
}
