package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	puzzle("test.txt")
	puzzle("input.txt")
}

type Equation struct {
	result   int
	operands []int
}

func parseInput(filepath string) []Equation {
	var equations []Equation 
	file, _ := os.Open(filepath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ": ")
		result, _ := strconv.Atoi(parts[0])
		values := strings.Split(parts[1], " ")
		operands := make([]int, len(values))
		for i, v := range values {
			o, _ := strconv.Atoi(v)
			operands[i] = o
		}
		equations = append(equations, Equation{result, operands})
	}
	return equations
}

func puzzle(path string) {
	part1:=0
	for _,eq := range  parseInput(path){
		if(hasValidEq(eq.result,eq.operands[0],eq.operands[1:])){
			part1 += eq.result
		}
	}
	fmt.Println(part1)
}

func hasValidEq(goal int, curr int,operands []int) bool{
	if(len(operands) == 0){
		return curr == goal
	}
	if(curr > goal){
		return false
	}
	conc,_ := strconv.Atoi(fmt.Sprintf("%d%d", curr, operands[0]))
	return hasValidEq(goal,curr+operands[0],operands[1:]) || 
		   hasValidEq(goal,curr*operands[0],operands[1:]) ||
		   hasValidEq(goal,conc,operands[1:]) 
}
