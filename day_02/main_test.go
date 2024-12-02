package main

import (
	"fmt"
	"testing"
)

func TestNumberOfSafeLevels(t *testing.T) {
	a := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	result := NumberOfSafeReports(a)
	fmt.Println(result)

	if result != 2 {
		t.Error("Expected 2, got", result)
	}
}

func TestNumberOfSafeLevelsErrorTolerant(t *testing.T) {
	a := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	result := NumberOfSafeLevelsErrorTolerant(a)
	fmt.Println(result)

	if result != 4 {
		t.Error("Expected 4, got", result)
	}
}

func TestNumberOfSafeLevelsErrorTolerantEdgeCases(t *testing.T) {
	a := `48 46 47 49 51 54 56
1 1 2 3 4 5
1 2 3 4 5 5
5 1 2 3 4 5
1 4 3 2 1
1 6 7 8 9
1 2 3 4 3
9 8 7 6 7
7 10 8 10 11
29 28 27 25 26 25 22 20`
	result := NumberOfSafeLevelsErrorTolerant(a)
	fmt.Println(result)

	if result != 10 {
		t.Error("Expected 10, got", result)
	}
}
