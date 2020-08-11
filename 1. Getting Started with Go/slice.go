package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var sli []int
	var sInput string

	sli = make([]int, 0, 3)

	for {
		fmt.Println("Enter a number:")
		fmt.Scanf("%s", &sInput)
		if sInput == "X" {
			break
		}

		iInput, _ := strconv.Atoi(sInput)
		sli = append(sli, iInput)
		sort.Ints(sli)
		fmt.Println(sli)
	}
}
