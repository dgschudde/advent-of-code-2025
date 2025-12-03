package main

import (
	"advent-of-code-2025/common"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input *[]string

	input = common.ReadInput("./input/input.txt")
	PartOne(*input)
}

func PartOne(input []string) {
	var total int

	for _, element := range input {
		// Convert element row to int array
		numbers := ConvertToIntArray(element)

		// Find highest number
		highest, highestIndex, highestOccurrence := FindMax(numbers, 9, 0)

		// Find second highest number
		secondHighest, secondHighestIndex, _ := FindMax(numbers, highest-1, highestIndex)

		var result strings.Builder

		if highestOccurrence > 1 {
			result.WriteString(strconv.Itoa(highest))
			result.WriteString(strconv.Itoa(highest))
			intResult, _ := strconv.Atoi(result.String())
			total += intResult
			continue
		}

		if highestIndex >= secondHighestIndex {
			result.WriteString(strconv.Itoa(secondHighest))
			result.WriteString(strconv.Itoa(highest))
		} else {
			result.WriteString(strconv.Itoa(highest))
			result.WriteString(strconv.Itoa(secondHighest))
		}

		intResult, _ := strconv.Atoi(result.String())
		total += intResult
	}
	fmt.Println(total)
}

func ConvertToIntArray(element string) []int {
	result := make([]int, len(element))
	for index, char := range element {
		result[index] = int(char - '0')
	}

	return result
}

func FindMax(numbers []int, maxToSearch int, startIndex int) (int, int, int) {
	maxValue := 0
	maxIndex := 0

	if startIndex+1 >= len(numbers) {
		startIndex = 0
	}

	for index := startIndex; index < len(numbers); index++ {
		if numbers[index] > maxValue && numbers[index] <= maxToSearch {
			maxValue = numbers[index]
			maxIndex = index
		}
	}

	occurrences := printUniqueValue(numbers)
	return maxValue, maxIndex, occurrences[maxValue]
}

func printUniqueValue(arr []int) map[int]int {
	dict := map[int]int{}
	for _, num := range arr {
		dict[num]++
	}
	return dict
}
