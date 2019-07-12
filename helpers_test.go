package main

import "testing"

func TestSumInts_hello(t *testing.T) {
	result := SumInts([]int{25, 25, 10})

	if result != 60 {
		t.Errorf("Solution incorrect, got: %d, want: %d", result, 60)
	}
}
