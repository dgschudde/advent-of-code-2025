package common

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Ranges struct {
	Min int
	Max int
}

func ReadInput(fileName string) *[]string {
	var input = make([]string, 0)

	// Read the input from a file
	inputFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(inputFile)

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return &input
}

func ConvertTo2DArray(input []string) [][]rune {
	dy := len(input)
	dx := len(input[0])

	arr := make([][]rune, dy)
	for i := range arr {
		arr[i] = make([]rune, dx)
	}

	for x, line := range input {
		for y, char := range line {
			arr[x][y] = char
		}
	}

	return arr
}

func ConvertToRanges(input []string) ([]Ranges, []int) {
	var ranges = make([]Ranges, 0)
	var ids = make([]int, 0)

	for _, line := range input {
		if strings.Contains(line, "-") {
			splitRange := strings.Split(line, "-")
			min, _ := strconv.Atoi(splitRange[0])
			max, _ := strconv.Atoi(splitRange[1])
			ranges = append(ranges, Ranges{min, max})
		} else if line == "" {
			continue
		} else {

			result, _ := strconv.ParseInt(line, 10, 64)
			ids = append(ids, int(result))
		}
	}

	return ranges, ids
}
