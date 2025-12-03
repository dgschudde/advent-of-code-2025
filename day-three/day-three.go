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
	PartTwo(*input)
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

func PartTwo(input []string) {
	var results int64
	var intResult int64

	for _, element := range input {
		result := Top12(element)

		intResult, _ = strconv.ParseInt(result, 10, 64)
		fmt.Println(intResult)
		results += intResult
	}
	fmt.Println(results)
}

func Top12(s string) string {
	n := len(s)
	k := 12
	if k >= n {
		return s
	}

	toRemove := n - k
	stack := make([]byte, 0, k)

	for i := 0; i < n; i++ {
		c := s[i]
		for toRemove > 0 && len(stack) > 0 && stack[len(stack)-1] < c {
			stack = stack[:len(stack)-1]
			toRemove--
		}
		stack = append(stack, c)
	}

	if toRemove > 0 {
		stack = stack[:len(stack)-toRemove]
	}

	return string(stack[:k])
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
