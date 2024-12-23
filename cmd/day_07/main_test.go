package main

import (
	"testing"
)

func TestSumValidEquationResults(t *testing.T) {
	testInput := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

	result := SumValidEquationResults(testInput)
	if result != 3749 {
		t.Errorf("Expected 3749, got %d", result)
	}
}

func TestSumValidEquationResultsPart2(t *testing.T) {
	testInput := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
3: 1 2 1000
100: 50 50 2`

	result := SumValidEquationResults2(testInput)
	if result != 11387 {
		t.Errorf("Expected 11387, got %d", result)
	}
}
