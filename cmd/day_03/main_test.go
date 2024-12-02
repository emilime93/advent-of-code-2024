package main

import "testing"

func TestMain(t *testing.T) {
	testLine := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	result := CalcMulSum(testLine)

	if result != 161 {
		t.Error("Expected 161, got ", result)
	}
}
