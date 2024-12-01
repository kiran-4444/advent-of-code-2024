package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"sort"
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
		part_1(input)

	}

	if part == 2 {
		part_2(input)

	}

}

func part_1(input string) {
	left_list, right_list := parseInput(input)

	sort.Ints(left_list)
	sort.Ints(right_list)

	ans := 0
	for index := range len(right_list) {
		ans += int(math.Abs(float64(left_list[index]) - float64(right_list[index])))
	}

	fmt.Println(ans)
}

func part_2(input string) {
	left_list, right_list := parseInput(input)

	ans := 0

	for _, left := range left_list {
		count := 0

		for _, right := range right_list {
			if right == left {
				count += 1
			}
		}

		ans += left * count
	}

	fmt.Println(ans)
}

func castToInt(input string) int {
	val, err := strconv.Atoi(input)
	if err != nil {
		panic("Error casting: " + err.Error())
	}

	return val

}

func parseInput(input string) ([]int, []int) {
	left_list := []int{}
	right_list := []int{}

	for _, group := range strings.Split(input, "\n") {
		values := strings.Fields(group)

		left_list = append(left_list, castToInt(values[0]))
		right_list = append(right_list, castToInt(values[1]))

	}

	return left_list, right_list

}
