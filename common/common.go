package common

import (
	"bufio"
	"log"
	"os"
)

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
