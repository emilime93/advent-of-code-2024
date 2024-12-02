package main

import (
	"adventofcode/pkg/input"
	"fmt"
	"strings"
)

func main() {
	input := input.ReadFile("./cmd/day_04/input.txt")
	numXMAS := CountXMAS(input)

	fmt.Printf("Found %d XMAXes\n", numXMAS)
}

func CountXMAS(input string) int {
	matrix := toMatrix(input)
	return findXMASes(matrix)
}

func findXMASes(matrix Matrix) int {
	xmases := 0
	for row := range matrix.rows {
		for col := range matrix.cols {
			xmases += hasXmas(row, col, matrix)
		}
	}
	return xmases
}

func hasXmas(row, col int, matrix Matrix) int {
	inBoundsRight := col+3 < matrix.cols
	inBoundsLeft := col > 2
	inBoundsUp := row > 2
	inBoundsDown := row+3 < matrix.rows
	numMatches := 0

	if inBoundsRight {
		if matrix.data[row][col] == "X" &&
			matrix.data[row][col+1] == "M" &&
			matrix.data[row][col+2] == "A" &&
			matrix.data[row][col+3] == "S" {
			numMatches++
		}
	}
	// left
	if inBoundsLeft {
		if matrix.data[row][col] == "X" &&
			matrix.data[row][col-1] == "M" &&
			matrix.data[row][col-2] == "A" &&
			matrix.data[row][col-3] == "S" {
			numMatches++
		}
	}
	// up
	if inBoundsUp {
		if matrix.data[row][col] == "X" &&
			matrix.data[row-1][col] == "M" &&
			matrix.data[row-2][col] == "A" &&
			matrix.data[row-3][col] == "S" {
			numMatches++
		}
	}
	// down
	if inBoundsDown {
		if matrix.data[row][col] == "X" &&
			matrix.data[row+1][col] == "M" &&
			matrix.data[row+2][col] == "A" &&
			matrix.data[row+3][col] == "S" {
			numMatches++
		}
	}
	// upright
	if inBoundsUp && inBoundsRight {
		if matrix.data[row][col] == "X" &&
			matrix.data[row-1][col+1] == "M" &&
			matrix.data[row-2][col+2] == "A" &&
			matrix.data[row-3][col+3] == "S" {
			numMatches++
		}
	}
	// downright
	if inBoundsDown && inBoundsRight {
		if matrix.data[row][col] == "X" &&
			matrix.data[row+1][col+1] == "M" &&
			matrix.data[row+2][col+2] == "A" &&
			matrix.data[row+3][col+3] == "S" {
			numMatches++
		}
	}
	// upleft
	if inBoundsUp && inBoundsLeft {
		if matrix.data[row][col] == "X" &&
			matrix.data[row-1][col-1] == "M" &&
			matrix.data[row-2][col-2] == "A" &&
			matrix.data[row-3][col-3] == "S" {
			numMatches++
		}
	}
	// downleft
	if inBoundsDown && inBoundsLeft {
		if matrix.data[row][col] == "X" &&
			matrix.data[row+1][col-1] == "M" &&
			matrix.data[row+2][col-2] == "A" &&
			matrix.data[row+3][col-3] == "S" {
			numMatches++
		}
	}
	return numMatches
}

func toMatrix(input string) Matrix {
	lines := strings.Split(input, "\n")
	numRows := len(lines)
	numCols := len(lines[0])
	matrix := make([][]string, numRows)
	for i, row := range lines {
		matrix[i] = strings.Split(row, "")
	}
	return Matrix{data: matrix, rows: numRows, cols: numCols}
}

func printMatrix(matrix Matrix) {
	for _, row := range matrix.data {
		for _, col := range row {
			fmt.Printf("%s", col)
		}
		fmt.Println()
	}
}

type Matrix struct {
	data [][]string
	rows int
	cols int
}
