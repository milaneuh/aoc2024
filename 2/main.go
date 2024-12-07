package main

import ( "bufio" "fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input [][]int = getInput()
	result := countValidReports(input)
	fmt.Println(result)
}

func countValidReports(reports [][]int) int {
    validCount := 0
    for _, report := range reports {
        if isReportValid(report) {
            validCount++
        }else{
			for index := range report{
				if(isReportValidWithoutIndex(report,index)){
					validCount++
					break
				}
			}
		}
    }
    return validCount
}

func isReportValid(report []int) bool {
	isIncreasing := report[0] < report[1]
    for i := 1; i < len(report); i++ {
        if isDifferenceNotInRange(report[i-1], report[i], isIncreasing) {
				return false
        }
    }
    return true
}

func isReportValidWithoutIndex(report []int, index int) bool {
	newReport := append([]int{}, report[:index]...)
    newReport = append(newReport, report[index+1:]...)	
	return isReportValid(newReport)
}

func isDifferenceNotInRange(previous int, current int, isIncreasing bool) bool {
    if isIncreasing != (previous < current) {
        return true
    }
    diff := abs(previous - current)
    return diff > 3 || diff < 1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func getInput() [][]int {
	var result [][]int
	file,err := os.Open("input.txt")

	if(err != nil){
		panic("Erreur d'ouverture du fichier")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		var report []int
		line := scanner.Text()
		parts := strings.Fields(line)

		for _,value := range parts {
			num,err := strconv.Atoi(value)
			if(err!=nil){
				panic("IoException during the conversion of the file")
			}
			report = append(report, num)
		}

		result = append(result,report)
	}
	return result
}
