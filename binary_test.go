package main

import "testing"

func TestBits_sum4294967295(t *testing.T) {
	result := SumInts(Bits())

	if result != 4294967295 {
		t.Errorf("Incorrect solution, got: %d, want: %d", result, 4294967295)
	}
}

func TestAsBinary_9is1001(t *testing.T) {
	result := AsBinary(9)

	if result != "1001" {
		t.Errorf("AsBinary(9) incorrect, got: %s, want: %s", result, "1001")
	}
}
