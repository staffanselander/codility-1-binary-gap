package main

import (
	"fmt"
	"testing"
)

func TestSolution_9_return2(t *testing.T) {
	answer := Solution(9)
	if answer != 2 {
		t.Errorf("Solution incorrect, got: %d, want: %d", answer, 2)
	}
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		input int
		want  int
	}{
		{9, 2},
		{32, 5},
		{18, 2},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%d=%d", testCase.input, testCase.want), func(t *testing.T) {
			got := Solution(testCase.input)

			if testCase.want != got {
				t.Errorf("want %d, got %d", testCase.want, got)
			}
		})
	}
}
