package main

import (
	"adventofcode/pkg/input"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	input := input.ReadFile("./cmd/day_03/input.txt")
	for _, line := range strings.Split(string(input), "\n") {
		sum += CalcMulSum(line)
	}
	fmt.Printf("Mul-sum: %d\n", sum)
}

func CalcMulSum(instructions string) int {
	regex := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")

	sum := 0
	matches := regex.FindAllStringSubmatch(instructions, -1)
	for _, submatches := range matches {
		var d1, d2 int
		d1 = toInt(submatches[1])
		d2 = toInt(submatches[2])
		sum += d1 * d2
	}
	return sum
}

func toInt(str string) int {
	integer, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("Could not convert \"%s\" into an int", str)
		panic(err)
	}
	return integer
}
