package main

import (
	"advent-of-code-2025/common"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input *[]string
	var splitInput []string

	input = common.ReadInput("./input/input.txt")

	for _, element := range *input {
		splitInput = strings.Split(element, ",")
	}

	totalPartOne := 0

	for _, element := range splitInput {
		splitElement := strings.Split(element, "-")
		left, _ := strconv.Atoi(splitElement[0])
		right, _ := strconv.Atoi(splitElement[1])

		totalPartOne += PartOne(left, right)
	}
	fmt.Println(totalPartOne)
}

func PartOne(left int, right int) int {
	result := 0
	for i := left; i <= right; i++ {
		if len(strconv.Itoa(i))%2 == 0 {
			e := strconv.Itoa(i)
			a := e[0 : len(e)/2]
			b := e[len(e)/2:]
			if a == b {
				result += i
			}
		}
	}
	return result
}
