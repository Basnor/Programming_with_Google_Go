package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sli []int
	var sInput string

	sli = make([]int, 0)

	fmt.Println("When you finish press X.")
	for i := 0; i < 10; i++ {
		fmt.Println("Enter a number:")
		fmt.Scanf("%s", &sInput)
		if sInput == "X" {
			break
		}

		iInput, _ := strconv.Atoi(sInput)
		sli = append(sli, iInput)

		bubbleSort(sli)
		fmt.Println(sli)
	}
}

func bubbleSort(items []int){
    for i := 0; i < len(items)-1; i++ {
		for j := 0; j < len(items)-i-1; j++ {
			if items[j] > items[j+1] {
				swap(items, j)
			}
		}
    }
}

func swap(items []int, i int) {
	items[i], items[i+1] = items[i+1], items[i]
}
