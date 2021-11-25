package main

import "testing"

func TestMatrix(t *testing.T) {
	want := [4][5]int{
		{130, 53, 21, 8, 4},
		{36, 16, 7, 3, 2},
		{8, 4, 2, 1, 1},
		{0, 0, 0, 0, 1},
	}
	got := Matrix()
	if got != want {
		t.Errorf("got = %v, want = %v", got, want)
	}
}
