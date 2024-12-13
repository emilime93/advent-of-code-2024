package main

import (
	"fmt"
	"testing"
)

func TestCountNumberOfVisitedPositions(t *testing.T) {
	fmt.Println("TEST")
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	result := CountNumberOfVisitedPositions(input)

	if result != 41 {
		t.Errorf("Expected 41, got %d", result)
	}
}
