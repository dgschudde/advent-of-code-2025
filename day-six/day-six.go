package main

import (
	"advent-of-code-2025/common"
	"fmt"
	"strconv"
	"strings"
)

var fileName string

func main() {
	var input *[]string

	fileName = "./input/input.txt"
	input = common.ReadInput(fileName)
	numbers, operators := ConvertCalculations(*input)

	total := PartOne(numbers, operators)

	fmt.Printf("Total: %d", total)
}

func PartOne(numbers [][]int, operators []string) int {
	size := len(numbers) - 1
	totalResult := 0

	for i := 0; i < len(operators); i++ {
		result := 0
		operator := operators[i]
		for j := 0; j < size; j++ {
			switch operator {
			case "+":
				result += numbers[j][i]
			case "*":
				if j == 0 {
					result = numbers[j][i] * numbers[j+1][i]
					j++
				} else {
					result *= numbers[j][i]
				}
			}
		}
		totalResult += result
	}

	return totalResult
}

func ConvertCalculations(input []string) ([][]int, []string) {
	items := make([]string, 0)
	arrNumbers := make([][]int, 0)
	arrOperators := make([]string, 0)

	for _, line := range input {
		items = strings.Split(line, " ")
		tempNumbers := make([]int, 0)
		for _, item := range items {
			if item != "" {
				if item == "*" || "+" == item {
					arrOperators = append(arrOperators, item)
				} else {
					itemToInt, _ := strconv.Atoi(item)
					tempNumbers = append(tempNumbers, itemToInt)
				}
			}
		}
		arrNumbers = append(arrNumbers, tempNumbers)
	}

	return arrNumbers, arrOperators
}
