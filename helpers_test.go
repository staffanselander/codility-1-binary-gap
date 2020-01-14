package main

import (
	"fmt"
	"testing"
)

func TestBits(t *testing.T) {
	want := []int{
		2147483648,
		1073741824,
		536870912,
		268435456,
		134217728,
		67108864,
		33554432,
		16777216,
		8388608,
		4194304,
		2097152,
		1048576,
		524288,
		262144,
		131072,
		65536,
		32768,
		16384,
		8192,
		4096,
		2048,
		1024,
		512,
		256,
		128,
		64,
		32,
		16,
		8,
		4,
		2,
		1,
	}
	got := Bits()

	for i, val := range want {
		if got[i] != val {
			t.Errorf("got %d, want %d", got[i], val)
		}
	}
}


func TestSumInts(t *testing.T) {
	testCases := []struct{
		input []int
		want  int
	}{
		{[]int{25,25,10}, 60},
		{[]int{10,10,10}, 30},
		{[]int{10,10,15}, 35},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%v=%d", testCase.input, testCase.want), func(t *testing.T) {
			got := SumInts(testCase.input)

			if testCase.want != got {
				t.Errorf("want %d, got %d", testCase.want, got)
			}
		})
	}
}
