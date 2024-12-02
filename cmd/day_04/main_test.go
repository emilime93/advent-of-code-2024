package main

import (
	"testing"
)

func TestCountXMAS(t *testing.T) {
	testInput := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	result := CountXMAS(testInput)

	if result != 18 {
		t.Error("Expected 18, got", result)
	}
}
