package main

import (
	"adventofcode/pkg/input"
	"fmt"
	"slices"
	"strings"
)

func main() {
	input := input.ReadFile("./cmd/day_06/input.txt")
	result := CountNumberOfVisitedPositions(input)
	fmt.Printf("Number of visited positions: %d\n", result)
}

func CountNumberOfVisitedPositions(input string) int {
	matrix := toMatrix(input)
	return simulate(matrix)
}

// The matrix has the following coordinate system:
// 0---------------------------> (x+) / cols
// |
// |
// |
// |
// |
// V
// (y+) / rows
func toMatrix(input string) Matrix {
	rows := strings.Split(input, "\n")
	numRows := len(rows)
	var guard Guard
	matrixData := make([][]rune, numRows)
	validDirections := []rune{'^', 'v', '>', '<'}
	for i, row := range rows {
		matrixData[i] = []rune(row)
		for j, char := range row {
			if slices.Contains(validDirections, char) {
				guard = Guard{
					pos:       Point{x: j, y: i},
					direction: directionFromRune(char),
				}
			}
		}
	}

	ySize := len(matrixData)
	xSize := len(matrixData[0])

	return Matrix{
		floorMap: matrixData,
		guard:    guard,
		xSize:    xSize,
		ySize:    ySize,
	}
}

func simulate(matrix Matrix) int {
	numberOfPositions := 0
	for {
		if matrix.guardIsExiting() {
			matrix.lastStep()
			numberOfPositions++
			break
		}
		if matrix.guardHasCollided() {
			matrix.turnGuard()
			continue
		}
		if matrix.step() {
			numberOfPositions++
		}
	}
	return numberOfPositions
}

func (m *Matrix) step() bool {
	x := m.guard.pos.x
	y := m.guard.pos.y
	var uniqueStep bool

	m.floorMap[y][x] = 'X'
	switch m.guard.direction {
	case North:
		m.guard.pos.y--
		if m.floorMap[y-1][x] == '.' {
			uniqueStep = true
		}
		m.floorMap[y-1][x] = '^'
	case East:
		m.guard.pos.x++
		if m.floorMap[y][x+1] == '.' {
			uniqueStep = true
		}
		m.floorMap[y][x+1] = '>'
	case South:
		m.guard.pos.y++
		if m.floorMap[y+1][x] == '.' {
			uniqueStep = true
		}
		m.floorMap[y+1][x] = 'V'
	case West:
		m.guard.pos.x--
		if m.floorMap[y][x-1] == '.' {
			uniqueStep = true
		}
		m.floorMap[y][x-1] = '<'
	default:
		panic(fmt.Sprintf("Unrecognized direction %d ", m.guard.direction))
	}
	return uniqueStep
}

func (m *Matrix) lastStep() {
	if !m.guardIsExiting() {
		panic("Cannot take last step if not exiting")
	}
	x := m.guard.pos.x
	y := m.guard.pos.y
	m.guard.pos = Point{x: -1, y: -1}
	m.floorMap[y][x] = 'X'
}

func (m *Matrix) guardHasCollided() bool {
	switch m.guard.direction {
	case North:
		if m.hasObstacleAt(m.guard.pos.x, m.guard.pos.y-1) {
			return true
		}
	case East:
		if m.hasObstacleAt(m.guard.pos.x+1, m.guard.pos.y) {
			return true
		}
	case South:
		if m.hasObstacleAt(m.guard.pos.x, m.guard.pos.y+1) {
			return true
		}
	case West:
		if m.hasObstacleAt(m.guard.pos.x-1, m.guard.pos.y) {
			return true
		}
	default:
		panic(fmt.Sprintf("unexpected main.Direction: %#v", m.guard.direction))
	}
	return false
}

func (m *Matrix) hasObstacleAt(x, y int) bool {
	if m.inBounds(x, y) {
		return m.floorMap[y][x] == '#'
	}
	return false
}

func (m *Matrix) inBounds(x, y int) bool {
	if (x < 0) ||
		(y < 0) ||
		(x >= m.xSize) ||
		(y >= m.ySize) {
		return false
	}
	return true
}

func (m *Matrix) turnGuard() {
	x := m.guard.pos.x
	y := m.guard.pos.y
	switch m.guard.direction {
	case North:
		m.guard.direction = East
		m.floorMap[y][x] = '>'
	case East:
		m.guard.direction = South
		m.floorMap[y][x] = 'V'
	case South:
		m.guard.direction = West
		m.floorMap[y][x] = '<'
	case West:
		m.guard.direction = North
		m.floorMap[y][x] = '^'
	}
}

func (m *Matrix) guardIsExiting() bool {
	guard := m.guard
	if (guard.pos.x == 0 && guard.direction == West) ||
		(guard.pos.y == 0 && guard.direction == North) ||
		(guard.pos.x == m.xSize-1 && guard.direction == East) ||
		(guard.pos.y == m.ySize-1 && guard.direction == South) {
		return true
	}
	return false
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Matrix struct {
	floorMap [][]rune
	guard    Guard
	xSize    int
	ySize    int
}

func (m *Matrix) setGuardAndDir(x int, y int, dir Direction) {
	var symbol rune
	switch dir {
	case North:
		symbol = '^'
	case East:
		symbol = '>'
	case South:
		symbol = 'V'
	case West:
		symbol = '<'
	}
	m.floorMap[y][x] = symbol
	m.guard.pos = Point{x: x, y: y}
}

func (m *Matrix) setGuard(x int, y int) {
	m.setGuardAndDir(x, y, m.guard.direction)
}

type Point struct {
	x int
	y int
}

type Guard struct {
	pos       Point
	direction Direction
}

func directionFromRune(r rune) Direction {
	switch r {
	case '^':
		return North
	case 'V':
		return South
	case '<':
		return West
	case '>':
		return East
	default:
		panic(fmt.Sprint("Unsupported direction: ", r))
	}
}

// for debug
func printMatrix(matrix Matrix) {
	for _, row := range matrix.floorMap {
		for _, col := range row {
			fmt.Printf("%s", string(col))
		}
		fmt.Println()
	}
}
