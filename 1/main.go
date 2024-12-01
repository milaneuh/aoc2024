package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	var l1, l2 []int
	
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Erreur d'ouverture du fichier:", err)
		return
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Format invalide dans la ligne:", line)
			continue
		}
	
		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Erreur de conversion des nombres:", line)
			continue
		}
	
		l1 = append(l1, num1)
		l2 = append(l2, num2)
	}

	//PART ONE
	l1 = mergeSort(l1)
	l2 = mergeSort(l2)

	i := 0
	r := 0
	for i < len(l1){
		if(l1[i] > l2[i]){
			r = r+ (l1[i] - l2[i])
		}else {
			r  = r+ (l2[i] - l1[i])
		}
		i++
	}

	fmt.Println(r)

	//PART TWO
	result := 0
	for _,location_id_l1 := range l1{
		multiplier := 0
		for _,location_id_l2 := range l2{
			if(location_id_l1 == location_id_l2){
				multiplier++
			}
		}
		result = result + (multiplier*location_id_l1)
	}
	fmt.Println(result)
}

func mergeSort(items []int) []int {
    if len(items) < 2 {
        return items
    }
    first := mergeSort(items[:len(items)/2])
    second := mergeSort(items[len(items)/2:])
    return merge(first, second)
}

func merge(a []int,	b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b){
		if a[i] < b[j] {
            final = append(final, a[i])
            i++
        } else {
            final = append(final, b[j])
            j++
        }
	}
    for ; i < len(a); i++ {
        final = append(final, a[i])
    }
    for ; j < len(b); j++ {
        final = append(final, b[j])
    }
    return final
}
