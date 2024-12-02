package main

import (
	"adventofcode/pkg/input"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := input.ReadFile("./cmd/day_02/input.txt")
	result := NumberOfSafeReports(input)

	fmt.Printf("Safe reports: %d\n", result)

	result2 := NumberOfSafeLevelsErrorTolerant(input)
	fmt.Printf("Safe levels with 1 skip: %d\n", result2)
}

func NumberOfSafeReports(input string) int {
	reports := strings.Split(input, "\n")
	safeReports := 0
	for _, report := range reports {
		if isSafeReport(report) {
			safeReports++
		}
	}
	return safeReports
}

func NumberOfSafeLevelsErrorTolerant(input string) int {
	reports := strings.Split(input, "\n")
	safeReports := 0
	for _, report := range reports {
		if isSafeReport(report) {
			safeReports++
		} else {
			levels := strings.Split(report, " ")
			for i := range levels {
				tmpLevels := make([]string, len(levels))
				copy(tmpLevels, levels)
				reportWithSkippedLevel := strings.Join(removeFromSlice(tmpLevels, i), " ")
				if isSafeReport(reportWithSkippedLevel) {
					safeReports++
					break
				}
			}
		}
	}
	return safeReports
}

func removeFromSlice(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func isSafeReport(level string) bool {
	levels := strings.Split(level, " ")
	increasing := false
	decreasing := false
	for i, level := range levels {
		if i == 0 {
			continue
		}
		prev := toInt(levels[i-1])
		current := toInt(level)
		// Decrease vs increase
		if prev > current {
			decreasing = true
		} else if prev < current {
			increasing = true
		}
		diff := int(math.Abs(float64(prev - current)))
		if diff > 3 || diff < 1 {
			return false
		}
		if increasing && decreasing {
			return false
		}
	}
	return true
}

func isSafeLevelWithError(level string) bool {
	values := strings.Split(level, " ")
	increasing := 0
	decreasing := 0
	faults := 0
	for i, value := range values {
		if i == 0 {
			continue
		}
		prev := toInt(values[i-1])
		current := toInt(value)
		// Decrease vs increase
		if prev > current {
			decreasing++
		} else if prev < current {
			increasing--
		}
		diff := int(math.Abs(float64(prev - current)))
		if diff > 3 || diff < 1 {
			faults++
		}
		if faults > 1 {
			faults++
		}
	}
	faults += int(math.Min(float64(increasing), float64(decreasing)))
	if faults > 1 {
		return false
	}
	return true
}

func toInt(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Could not convert \"%s\" into an int", value)
		panic(err)
	}
	return intValue
}
