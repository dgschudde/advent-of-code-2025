package main

import (
	"advent-of-code-2025/common"
	"fmt"
)

type Status int

const (
	Free Status = iota
	Taken
	OutOfBounds
)

type Position struct {
	x int
	y int
}

func main() {
	var input *[]string

	input = common.ReadInput("./input/input.txt")
	arr2D := common.ConvertTo2DArray(*input)

	PartOne(&arr2D)

	totalRemoved := 0

	for {
		total := PartTwo(&arr2D)
		totalRemoved += total
		if total == 0 {
			break
		}
	}
	fmt.Printf("Total removed %d\r\n", totalRemoved)
}

func PartOne(input *[][]rune) {
	totalCount := 0

	for y, line := range *input {
		for x, char := range line {
			if char == '@' {
				//fmt.Printf("x:%d Y:%d | ", x, y)
				total := CheckAllDirections(x, y, input)
				if total < 4 {
					fmt.Printf("x")
					totalCount++
				} else {
					fmt.Printf("%c", char)
				}
			} else {
				fmt.Printf("%c", char)
			}
		}
		fmt.Println()
	}

	fmt.Printf("Total free %d\r\n", totalCount)
}

func PartTwo(input *[][]rune) int {
	totalCount := 0

	arr := *input
	positions := make([]Position, 0)

	for y, line := range arr {
		for x, char := range line {
			if char == '@' {
				total := CheckAllDirections(x, y, input)
				if total < 4 {
					positions = append(positions, Position{x, y})
					totalCount++
				}
			}
		}
	}

	// Remove all positions from the list
	for _, position := range positions {
		arr[position.y][position.x] = '.'
	}

	return totalCount
}

func CheckAllDirections(x int, y int, input *[][]rune) int {
	var status Status
	var result = 0

	status = CheckNorth(x, y, input)
	if status == Taken {
		result++
	}

	status = CheckSouth(x, y, input)
	if status == Taken {
		result++
	}

	status = CheckWest(x, y, input)
	if status == Taken {
		result++
	}

	status = CheckEast(x, y, input)
	if status == Taken {
		result++
	}

	status = CheckNorthEast(x, y, input)
	if status == Taken {
		result++
	}

	status = CheckNorthWest(x, y, input)
	if status == Taken {
		result++
	}
	status = CheckSouthEast(x, y, input)
	if status == Taken {
		result++
	}
	status = CheckSouthWest(x, y, input)
	if status == Taken {
		result++
	}

	return result
}

func CheckNorth(x int, y int, input *[][]rune) Status {
	arr := *input

	if y <= 0 {
		return OutOfBounds
	}

	if arr[y-1][x] == '@' {
		return Taken
	}
	return Free
}

func CheckSouth(x int, y int, input *[][]rune) Status {
	arr := *input

	if y >= len(arr)-1 {
		return OutOfBounds
	}

	if arr[y+1][x] == '@' {
		return Taken
	}
	return Free
}

func CheckWest(x int, y int, input *[][]rune) Status {
	arr := *input

	if x <= 0 {
		return OutOfBounds
	}

	if arr[y][x-1] == '@' {
		return Taken
	}
	return Free
}

func CheckEast(x int, y int, input *[][]rune) Status {
	arr := *input

	if x >= len(arr[0])-1 {
		return OutOfBounds
	}

	if arr[y][x+1] == '@' {
		return Taken
	}
	return Free
}

func CheckNorthEast(x int, y int, input *[][]rune) Status {
	arr := *input

	if y <= 0 || x >= len(arr[0])-1 {
		return OutOfBounds
	}

	if arr[y-1][x+1] == '@' {
		return Taken
	}
	return Free
}

func CheckNorthWest(x int, y int, input *[][]rune) Status {
	arr := *input

	if y <= 0 || x <= 0 {
		return OutOfBounds
	}

	if arr[y-1][x-1] == '@' {
		return Taken
	}
	return Free
}

func CheckSouthEast(x int, y int, input *[][]rune) Status {
	arr := *input

	if y >= len(arr)-1 || x >= len(arr[0])-1 {
		return OutOfBounds
	}

	if arr[y+1][x+1] == '@' {
		return Taken
	}
	return Free
}

func CheckSouthWest(x int, y int, input *[][]rune) Status {
	arr := *input

	if y >= len(arr)-1 || x <= 0 {
		return OutOfBounds
	}

	if arr[y+1][x-1] == '@' {
		return Taken
	}
	return Free
}
