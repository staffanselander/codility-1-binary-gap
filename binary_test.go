package main

import (
	"fmt"
	"testing"
)

func TestAsBinary(t *testing.T) {
	testCases := []struct {
		input int
		want  string
	}{
		{9, "1001"},
		{32, "100000"},
		{18, "10010"},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%d = %s", testCase.input, testCase.want), func(t *testing.T) {
			got := AsBinary(testCase.input)

			if testCase.want != got {
				t.Errorf("want %s, got %s", testCase.want, got)
			}
		})
	}
}
