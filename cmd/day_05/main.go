package main

import (
	"adventofcode/pkg/input"
	"adventofcode/pkg/util"
	"fmt"
	"math"
	"slices"
	"strings"
)

func main() {
	input := input.ReadFile("./cmd/day_05/input.txt")
	resutl := SumOfCorrectMiddlePages(input)
	fmt.Printf("Sum: %d\n", resutl)
}

func SumOfCorrectMiddlePages(input string) int {
	sum := 0
	parts := strings.Split(input, "\n\n")
	rules := getRules(parts[0])
	updates := getUpdates(parts[1])

	for n, update := range updates {
		middle, err := isCorrectUpdate(update, rules)
		if err != nil {
			fmt.Printf("%d ERROR: %v\n", n, err)
			continue
		}
		sum += middle
	}
	return sum
}

const EMPTY int = 0

func isCorrectUpdate(update []int, rules map[int][]int) (int, error) {
	for i, page := range update {
		mustBePrecededByPages, err := pagesCalc(page, rules[page], update)
		if err != nil {
			fmt.Printf("Page %d did not have any applicable/present rules\n", page)
			continue
		}

		fmt.Printf("%d must be preceeded by %v\n", page, mustBePrecededByPages)
		for _, precededByPage := range mustBePrecededByPages {
			indexOfPreceded := slices.Index(update, precededByPage)
			if indexOfPreceded >= i {
				return 0, fmt.Errorf("%d was not preceded by %v (indexOfPreceded=%d)\n%v", page, mustBePrecededByPages, indexOfPreceded, update)
			}
			fmt.Printf("%d was preceeded by %d @%d", page, precededByPage, indexOfPreceded)
		}
	}
	middleIndex := int(math.Floor(float64(len(update)) / float64(2)))
	fmt.Printf("Returning %d\n", update[middleIndex])
	return update[middleIndex], nil
}

func pagesCalc(page int, precededByPages []int, update []int) ([]int, error) {
	if precededByPages == nil {
		fmt.Printf("%d does not have a rule, counts as OK\n", page)
		return nil, fmt.Errorf("Page %d did not have any associated rules", page)
	}
	precededByPagesPresent, err := contains(update, precededByPages)
	if err != nil {
		return nil, err
	}
	return precededByPagesPresent, nil
}

func getUpdates(updatesString string) [][]int {
	size := len(strings.Split(updatesString, "\n"))
	updates := make([][]int, size)
	for i, update := range strings.Split(updatesString, "\n") {
		updatePagesStrings := strings.Split(update, ",")
		updatePages := make([]int, len(updatePagesStrings))
		for j, updatePage := range updatePagesStrings {
			updatePages[j] = util.ToInt(updatePage)
		}
		updates[i] = updatePages
	}
	return updates
}

// A rule [4]=5 means that 4 must be preceeded by 5, if 5 exists
// "Key must be preceeded by value"
func getRules(rulesString string) map[int][]int {
	rules := make(map[int][]int)
	rulesRows := strings.Split(rulesString, "\n")
	for _, ruleRow := range rulesRows {
		rulesSplit := strings.Split(ruleRow, "|")
		page := util.ToInt(rulesSplit[1])
		precededBy := util.ToInt(rulesSplit[0])
		if rules[page] == nil {
			rules[page] = make([]int, 0)
		}
		rules[page] = append(rules[page], precededBy)
	}
	fmt.Printf("Rules: %v\n", rules)
	return rules
}

func contains(source []int, compare []int) ([]int, error) {
	contained := make([]int, 0)
	for _, item := range compare {
		if slices.Contains(source, item) {
			contained = append(contained, item)
		}
	}
	if len(contained) == 0 {
		return nil, fmt.Errorf("No occurances of %v in %v", compare, source)
	}
	return contained, nil
}
