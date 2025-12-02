package main

import (
	"advent-of-code-2025/common"
	"fmt"
	"strconv"
)

var outComes map[int]int

func main() {
	var input *[]string

	input = common.ReadInput("./input/input.txt")

	partOne(*input)
	partTwo(*input)
}

func partOne(input []string) {
	var dial = 50

	outComes = make(map[int]int, len(input))

	for index, element := range input {
		var direction = string(element[0])
		var distance, _ = strconv.Atoi(element[1:])

		distance = distance % 100

		switch direction {
		case "R":
			{
				dial += distance
				if dial > 99 {
					dial = dial - 100
				}
			}
		case "L":
			{
				dial -= distance
				if dial < 0 {
					dial = dial + 100
				}
			}
		}
		//fmt.Printf("Direction: %s => %d dial: %d\r\n", direction, distance, dial)
		outComes[index] = dial
	}

	var occurrences = 0
	for _, number := range outComes {
		if number == 0 {
			occurrences++
		}
	}

	fmt.Printf("Password: %d\r\n", occurrences)
}

func partTwo(input []string) {
	var dial = 50
	var bounds = 0

	outComes = make(map[int]int, len(input))

	for index, element := range input {
		var direction = string(element[0])
		var distance, _ = strconv.Atoi(element[1:])

		bounds += distance / 100
		distance = distance % 100

		switch direction {
		case "R":
			{
				dial += distance
				if dial > 99 {
					dial = dial - 100
					bounds++
				}
			}
		case "L":
			{
				dial -= distance
				if dial < 0 {
					dial = dial + 100
					bounds++
				}
			}
		}
		//fmt.Printf("Direction: %s => %d dial: %d\r\n", direction, distance, dial)
		if dial == 0 && distance/100 > 0 {
			//bounds--
		}
		outComes[index] = dial
	}

	var occurrences = 0
	for _, number := range outComes {
		if number == 0 {
			occurrences++
		}
	}

	fmt.Printf("Password: %d\r\n", bounds)
}
