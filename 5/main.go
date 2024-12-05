package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Pair struct {
	left int
	right int
}

func main() {
	filePath := "input.txt"

	pairs, matrix, _ := ParseFile(filePath)
	part1 := 0
	part2 := 0

	for _,row := range matrix {
	 	reOrderedRow := reOrderRow(row,pairs)
		if slices.Equal(reOrderedRow,row) {
			part1 += row[(len(row)/2)]
		}else {
			part2 += reOrderedRow[(len(reOrderedRow)/2)]
		}
	} 
	
	fmt.Println(part1)
	fmt.Println(part2)
}

func reOrderRow(row []int, pairs []Pair) []int {
	flag := true
	reOrderedRow := slices.Clone(row)
	for index,num := range row {
		for _,pair :=  range pairs {
			if(pair.left == num){
				rightIndex := slices.Index(row,int(pair.right))
				if(flag && rightIndex != -1 && rightIndex <  index){
					flag = false
					reOrderedRow[rightIndex] = num
					reOrderedRow[index] = pair.right
				}
			}else if(pair.right == num){
				leftIndex := slices.Index(row,int(pair.left))
				if(flag && leftIndex != -1 && leftIndex >  index){
					flag = false
					reOrderedRow[leftIndex] = num
					reOrderedRow[index] = pair.left
				}
			}
		}
	}
	if(!flag){
		return reOrderRow(reOrderedRow,pairs)
	}else {
		return reOrderedRow
	}
}

func ParseFile(filePath string) ([]Pair, [][]int, error) {
	file, _ := os.Open(filePath)
	defer file.Close()
	var pairs []Pair
	var matrix [][]int
	scanner := bufio.NewScanner(file)
	isMatrix := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// If the line is empty, switch to matrix parsing
		if line == "" {
			isMatrix = true
			continue
		}

		if !isMatrix {
			// Parse pairs
			parts := strings.Split(line, "|")
			l, _ := strconv.Atoi(parts[0])
			r, _ := strconv.Atoi(parts[1])
			pairs = append(pairs, Pair{left: l, right: r})
		} else {
			// Parse matrix
			values := strings.Split(line, ",")
			var row []int
			for _, val := range values {
				num, _ := strconv.Atoi(strings.TrimSpace(val))
				row = append(row, num)
			}
			matrix = append(matrix, row)
		}
	}
	return pairs, matrix, nil
}