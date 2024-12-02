package main

import (
	"fmt"
	"testing"
)

func TestCalculateDistanceSum(t *testing.T) {
	a := `3   4
	4   3
	2   5
	1   3
	3   9
	3   3`
	result := CalculateDistanceSum(a)
	fmt.Println(result)

	if result != 11 {
		t.Error("Expected 11")
	}

}

func TestCalculateSimilarity(t *testing.T) {
	a := `3   4
	4   3
	2   5
	1   3
	3   9
	3   3`
	result := CalculateSimilarity(a)
	fmt.Println(result)

	if result != 31 {
		t.Error("Expected 31")
	}

}
