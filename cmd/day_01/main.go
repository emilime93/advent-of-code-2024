package main

import (
	"adventofcode/pkg/input"
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := input.ReadFile("./cmd/day_01/input.txt")
	// Part 1
	distances := CalculateDistanceSum(input)
	fmt.Printf("Distances: %d\n", distances)

	// Part 2
	similarityScore := CalculateSimilarity(input)
	fmt.Printf("Similarity: %d\n", similarityScore)
}

func CalculateSimilarity(input string) int {
	left, right := toLists(input)
	similarities := make([]int, 0, 1000)
	for _, l := range left {
		occurances := numberOfOccurances(l, right)
		similarity := l * occurances
		similarities = append(similarities, similarity)
	}
	sum := 0
	for _, subSum := range similarities {
		sum += subSum
	}
	return sum
}

func numberOfOccurances(l int, list []int) int {
	count := 0
	for _, r := range list {
		if l == r {
			count++
		}
	}
	return count
}

func CalculateDistanceSum(input string) int {
	left, right := toLists(input)
	slices.Sort(left)
	slices.Sort(right)
	pairs := toPairs(left, right)
	distancesSum := sumDistances(pairs)
	return distancesSum
}

func sumDistances(pairs []Pair) int {
	sum := 0
	for _, pair := range pairs {
		sum += pair.distance()
	}
	return sum
}

func toPairs(left []int, right []int) []Pair {
	pairs := make([]Pair, 0, len(left))
	for i := range left {
		pairs = append(pairs, Pair{left: left[i], right: right[i]})
	}
	return pairs
}

const TRIPPLE_SPACE = "   "

func toLists(input string) ([]int, []int) {
	// var left []string
	left := make([]int, 0, 1000)
	right := make([]int, 0, 1000)

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		split := strings.Split(line, TRIPPLE_SPACE)
		leftInt := toInt(split[0])
		rightInt := toInt(split[1])

		left = append(left, leftInt)
		right = append(right, rightInt)
	}
	return left, right
}

func toInt(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Could not convert %s into an int", value)
		panic(err)
	}
	return intValue
}

type Pair struct {
	left  int
	right int
}

func (p *Pair) distance() int {
	distance := p.left - p.right
	return int(math.Abs(float64(distance)))
}
