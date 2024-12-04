package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(part1("input.txt"))
	fmt.Println(part2("input.txt"))
}

func part2(path string) int {
	xmas := 0
	puzzle := getPuzzle(path)
	for row := 1; row < len(puzzle) - 1 ; row ++ {
		for col := 1; col < len(puzzle[row]) -1; col++ {
				if 
					string(puzzle[row][col]) == "A" &&
				//  ↖									       ↘
					((string(puzzle[row - 1][col - 1]) == "S" && string(puzzle[row + 1][col + 1]) == "M") ||
					(string(puzzle[row - 1][col - 1]) == "M" && string(puzzle[row + 1][col + 1]) == "S")) &&
				//  ↙									       ↗
					((string(puzzle[row + 1][col - 1]) == "S" && string(puzzle[row - 1][col + 1]) == "M") ||
					(string(puzzle[row + 1][col - 1]) == "M" && string(puzzle[row - 1][col + 1]) == "S")) {
						xmas += 1
					}
		}
	}
	return xmas
}

func part1(path string) int{
	return mapPuzzleDirections(getPuzzle(path))
}

func mapPuzzleDirections(puzzle []string ) int{
	var result int

	for lineIndex,line := range puzzle {
		for index := range line {
			result += part1Puzzle(puzzle,lineIndex,index,line)
		}
	}
	return result
}

func part1Puzzle(puzzle []string, lineIndex int, index int , line string) int{
	var result int
	if len(line) - index >= 4 && isXmasRightHorizontal(line,index){
		result += 1
	}
	if index >= 3 && isXmasLeftHorizontal(line,index) {
		result += 1
	}
	if len(puzzle) - lineIndex > 3 {
		if isXmasDownVertical(puzzle,lineIndex,index){
			result += 1
		}
		if index >= 3  && isXmasLeftDiagonal(puzzle,lineIndex,index,true){
			result += 1
		}
		if len(line) - index >= 4 && isXmasRightDiagonal(puzzle,lineIndex,index,true){
			result += 1
		}
	}
	if lineIndex >= 3  {
		if isXmasUpVertical(puzzle,lineIndex,index){
			result += 1
		}
		if index >= 3  && isXmasLeftDiagonal(puzzle,lineIndex,index,false){
			result += 1
		}
		if len(line) - index >= 4 && isXmasRightDiagonal(puzzle,lineIndex,index,false){
			result += 1
		}
	}
	return result
}

func isXmasLeftDiagonal(puzzle []string , lineIndex int, index int, downDirection bool) bool{
	var result string
	if(downDirection){
		for i := lineIndex; i < lineIndex+4; i++{
			result = result + string(puzzle[i][index])
			index = index - 1
		}
	}else {
		for i := lineIndex; i > lineIndex-4; i--{
			result = result + string(puzzle[i][index])
			index = index - 1
		}
	}
	return result == "XMAS"
}

func isXmasRightDiagonal(puzzle []string , lineIndex int, index int, downDirection bool) bool{
	var result string
	if(downDirection){
		for i := lineIndex; i < lineIndex+4; i++{
			result = result + string(puzzle[i][index])
			index = index + 1
		}
	}else {
		for i := lineIndex; i > lineIndex-4; i--{
			result = result + string(puzzle[i][index])
			index = index + 1
		}
	}
	return result == "XMAS"
}

func isXmasUpVertical(puzzle []string , lineIndex int, index int) bool{
	var result string
	for i := lineIndex; i > lineIndex-4; i--{
		result = result + string(puzzle[i][index])
	}
	return result == "XMAS"
}

func isXmasDownVertical(puzzle []string , lineIndex int, index int) bool{
	var result string
	for i := lineIndex; i < lineIndex+4; i++{
		result = result + string(puzzle[i][index])
	}
	return result == "XMAS"
}

func isXmasLeftHorizontal(line string, index int) bool{
	var result string
	for i := index; i > index-4; i --{
		result = result + string(line[i])
	}
	return result == "XMAS"
}

func isXmasRightHorizontal(line string, index int) bool{
	var result string
	for i := index; i < index+4; i ++{
		result = result + string(line[i])
	}
	return result == "XMAS"
}

func getPuzzle(path string) []string {
	var puzzle []string
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan(){
		puzzle = append(puzzle, scanner.Text())
	}
	return puzzle
}

