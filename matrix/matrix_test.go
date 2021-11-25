package main

import (
	"reflect"
	"testing"
)

func TestGen(t *testing.T) {
	for _, tc := range []struct {
		width, height int
		want          [][]int
	}{
		{0, 0, [][]int{}},
		{1, 1, [][]int{
			{0},
		}},
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
func equal(a, b [][]int) bool {
	return reflect.DeepEqual(a, b)
}
