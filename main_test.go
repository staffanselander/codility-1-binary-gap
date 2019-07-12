package main

import (
	"testing"
)

func TestSolution_9_return2(t *testing.T) {
	answer := Solution(9)
	if answer != 2 {
		t.Errorf("Solution incorrect, got: %d, want: %d", answer, 2)
	}
}
