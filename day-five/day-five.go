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

	input = common.ReadInput("./input/test-input.txt")
	ranges, ids := common.ConvertToRanges(*input)
	PartOne(ranges, ids)
}

func PartOne(ranges []common.Ranges, ids []int) {
	var validIds = make([]int, 0)

	for _, id := range ids {
		//fmt.Println(id)
		//fmt.Println()
		for _, item := range ranges {
			//fmt.Println(item)
			if id >= item.Min && id <= item.Max {
				validIds = append(validIds, id)
				break
			}
		}
		//fmt.Println()
	}

	//fmt.Println(validIds)
	fmt.Println(len(validIds))
}
