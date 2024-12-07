package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n")
	var matrix [][]string
	row := 0
	index := 0
	for i, val := range split {
		var line []string
		for _, s := range val {
			if string(s) == "^" {
				index = len(line)
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
		newRowIndex := row + directions[direction][0]
		newColIndex := index + directions[direction][1]
		newVal, err := getNewIndex(newRowIndex, newColIndex, matrix)
		// Check if next jump is outOfBounds
		if err != nil {
			break
		}
		// Check if next jump has obstacle
		if newVal == "#" {
			if direction == 3 {
				direction = 0
			} else {
				direction += 1
			}
			continue
		}
		fmt.Println("Current : ", row, " ", index, " New : ", newRowIndex, " ", newColIndex)
		// Remplace current tile with X if it's a dot
		if string(matrix[row][index]) == "." {
			matrix[row][index] = "X"
			part1 += 1
		}
		row = newRowIndex
		index = newColIndex
	}
	fmt.Println(part1)
}

func getNewIndex(row int, col int, matrix [][]string) (string, error) {
	if row < 0 || row > len(matrix)-1 || col < 0 || col > len(matrix[0])-1 {
		return "", errors.New("Out of bounds")
	} else {
		return string(matrix[row][col]), nil
	}
}
