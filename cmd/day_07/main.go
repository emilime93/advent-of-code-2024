package main

import (
	"adventofcode/pkg/input"
	"adventofcode/pkg/util"
	"fmt"
	"strings"
)

func main() {
	input := input.ReadFile("cmd/day_07/input.txt")
	result1 := SumValidEquationResults(input)
	result2 := SumValidEquationResults2(input)
	fmt.Printf("Sum of valid equations: %d\n", result1)
	fmt.Printf("Sum of valid equations (part 2): %d\n", result2)
}

func SumValidEquationResults2(input string) int {
	equations := inputToEquations(input)
	sum := 0
	for _, equation := range equations {
		if equation.isValidWithConcatenation() {
			sum += equation.result
		}
	}
	return sum
}

func SumValidEquationResults(input string) int {
	equations := inputToEquations(input)
	sum := 0
	for _, equation := range equations {
		if equation.isValid() {
			sum += equation.result
		}
	}
	return sum
}

func inputToEquations(input string) []Equation {
	equations := make([]Equation, 0)
	for _, line := range strings.Split(input, "\n") {
		lineParts := strings.Split(line, ": ")
		result := util.ToInt(lineParts[0])
		operands := make([]int, 0)
		for _, operand := range strings.Split(lineParts[1], " ") {
			operands = append(operands, util.ToInt(operand))
		}
		equation := Equation{result: result, operands: operands}
		equations = append(equations, equation)
	}
	return equations
}

type Equation struct {
	result   int
	operands []int
}

func (e *Equation) isValid() bool {
	return recursiveIsValid(e.result, e.operands, 0)
}

func recursiveIsValid(goalResult int, operands []int, result int) bool {
	if len(operands) == 0 {
		return goalResult == result
	}
	addResult := result + operands[0]
	// You may not "cancel" the first number by multiplying it with 0...
	// This is an effect of me starting the reccursion with 0
	var mulResult int
	if result == 0 {
		mulResult = operands[0]
	} else {
		mulResult = result * operands[0]
	}
	return recursiveIsValid(goalResult, operands[1:], addResult) ||
		recursiveIsValid(goalResult, operands[1:], mulResult)
}

func (e *Equation) isValidWithConcatenation() bool {
	return recursiveIsValidWithConcatenation(e.result, e.operands, 0, 0)
}

func recursiveIsValidWithConcatenation(goalResult int, operands []int, result int, depth int) bool {
	if len(operands) == 0 {
		return goalResult == result
	}
	addResult := result + operands[0]
	// You may not "cancel" the first number by multiplying it with 0...
	// This is an effect of me starting the reccursion with 0
	var mulResult int
	if result == 0 {
		mulResult = operands[0]
	} else {
		mulResult = result * operands[0]
	}
	concatResult := util.ToInt(fmt.Sprintf("%d%d", result, operands[0]))
	return recursiveIsValidWithConcatenation(goalResult, operands[1:], addResult, depth+1) ||
		recursiveIsValidWithConcatenation(goalResult, operands[1:], mulResult, depth+1) ||
		recursiveIsValidWithConcatenation(goalResult, operands[1:], concatResult, depth+1)
}
