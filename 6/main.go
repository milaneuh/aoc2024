package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("test.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n")
	var matrix [][]string
	row := 0
	index := 0
	for i, val := range split {
		var line []string
		for _, s := range val {
			if string(s) == "^" {
				index = len(line) - 1
				row = i
			}
			line = append(line, string(s))
		}
		matrix = append(matrix, line)
	}

	directions := [][]int{
		{-1, 0}, // top
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}

	part1 := 2 // 2 because of start and end
	direction := 0
	for {
		// Check if next jump is outOfBounds
		if (row == 0 && direction == 0) ||
			(row == len(matrix[row]) && direction == 2) ||
			(index == 0 && direction == 3) ||
			(index == len(matrix[row])-1 && direction == 1) {
			break
		}
		// Check if next jump has obstacle
		fmt.Println(direction, " ", row, " ", index, " ", directions[direction], " ", len(matrix[row]), " ", len(matrix))
		if string(matrix[row+directions[direction][0]][index+directions[direction][1]]) == "#" {
			if direction == 3 {
				direction = 0
			} else {
				direction += 1
			}
		}
		// Remplace current tile with X if it's a dot
		if string(matrix[row][index]) == "." {
			matrix[row][index] = "X"
			part1 += 1
		}
		row = row + directions[direction][0]
		index = index + directions[direction][1]
	}
	fmt.Println(part1)
}
