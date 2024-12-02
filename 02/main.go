package main

import (
	"cmp"
	_ "embed"
	"flag"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("Empty content!")
	}

	var part int
	flag.IntVar(&part, "part", 1, "part 1 or part 2")
	flag.Parse()
	fmt.Println("Running part: ", part)

	if part == 1 {
		part1(input)
	}

	if part == 2 {
		part2(input)

	}
}

func part1(input string) {
	parsedInput := parseInput(input)

	ans := 0

	for _, group := range parsedInput {
		if (slices.IsSorted(group) || slices.IsSortedFunc(group, func(a, b int) int {
			return cmp.Compare(b, a)
		})) && checkDiff(group) {
			ans += 1
		}
	}

	fmt.Println(ans)

}

func part2(input string) {
	parsedInput := parseInput(input)
	ans := 0
	for _, group := range parsedInput {
		if (slices.IsSorted(group) || slices.IsSortedFunc(group, func(a, b int) int {
			return cmp.Compare(b, a)
		})) && checkDiff(group) {
			ans += 1
		} else {
			for i := range len(group) {
				removedSlice := remove(group, i)
				if (slices.IsSorted(removedSlice) || slices.IsSortedFunc(removedSlice, func(a, b int) int {
					return cmp.Compare(b, a)
				})) && checkDiff(removedSlice) {
					ans += 1
					break
				}
			}
		}
	}
	fmt.Println(ans)

}

func remove(slice []int, index int) []int {
	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)
	return append(sliceCopy[:index], sliceCopy[index+1:]...)
}

func checkDiff(innerList []int) bool {
	// fmt.Println("innerList: ", innerList)
	for i := range len(innerList) {
		if i == 0 {
			continue
		}

		currDiff := math.Abs(float64(innerList[i] - innerList[i-1]))
		// fmt.Println("currDiff: ", currDiff)

		if currDiff != 1 && currDiff != 2 && currDiff != 3 {
			return false
		}

	}

	return true
}

func castToInt(input string) int {
	val, err := strconv.Atoi(input)
	if err != nil {
		panic("Error casting: " + err.Error())
	}

	return val
}

func parseInput(input string) [][]int {
	outerList := [][]int{}

	for _, group := range strings.Split(input, "\n") {
		values := strings.Fields(group)

		innerList := []int{}

		for _, value := range values {
			innerList = append(innerList, castToInt(value))
		}
		outerList = append(outerList, innerList)
	}

	return outerList

}
