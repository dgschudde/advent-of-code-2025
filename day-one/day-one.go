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

	//partOne(*input)
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
	var countZeroes = 0
	var previousDial = -1

	for _, element := range input {
		var direction = string(element[0])
		var distance, _ = strconv.Atoi(element[1:])

		hundreds := distance / 100
		distance = distance % 100
		countZeroes += hundreds

		switch direction {
		case "R":
			{
				dial += distance
				if dial > 99 {
					dial = dial - 100
					countZeroes++
					if previousDial == 0 {
						countZeroes++
					}
				}
			}
		case "L":
			{
				dial -= distance
				if dial == 0 {
					countZeroes++
					if previousDial == 0 {
						countZeroes++
					}
				}
				if dial < 0 {
					dial = dial + 100
					if previousDial == 0 {
						countZeroes++
					}
				}
			}
		}

		previousDial = dial

		fmt.Printf("Direction: %s => %d dial: %d\r\n", direction, distance, dial)
	}

	fmt.Printf("Password: %d\r\n", countZeroes)
}
