package main

import (
	"testing"
)

func TestFun(t *testing.T) {
	in := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	out := findCircleNum(in)
	want := 2
	if out != want {
		t.Errorf("got %d, want %d", out, want)
	}
}
