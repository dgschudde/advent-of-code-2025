package main

import (
	"advent-of-code-2025/common"
	"fmt"
	"strings"
)

var fileName string

func main() {
	var input *[]string

	fileName = "./input/input.txt"
	input = common.ReadInput(fileName)
	servers := ParseInput(input)
	PartOne(servers)
}

func ParseInput(input *[]string) map[string][]string {
	servers := make(map[string][]string)
	// Split "aaa: you hhh"
	for _, line := range *input {
		parts := strings.SplitN(line, ":", 2)

		input := strings.TrimSpace(parts[0])
		outputsText := strings.TrimSpace(parts[1])

		var outputs []string
		if outputsText != "" {
			outputs = strings.Fields(outputsText) // split by space
		}

		servers[input] = outputs
	}
	return servers
}

func PartOne(graph map[string][]string) {
	paths := findAllPaths(graph, "you", "out")
	fmt.Println(len(paths))
}

func findAllPaths(graph map[string][]string, start, end string) [][]string {
	var result [][]string
	var path []string

	var dfs func(node string)
	dfs = func(node string) {
		path = append(path, node)

		// Reached destination
		if node == end {
			// Copy current path
			cp := make([]string, len(path))
			copy(cp, path)
			result = append(result, cp)
			path = path[:len(path)-1]
			return
		}

		for _, next := range graph[node] {
			dfs(next)
		}

		// backtrack
		path = path[:len(path)-1]
	}

	dfs(start)
	return result
}
