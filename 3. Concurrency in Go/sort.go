package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var sli []int
	sli = readInt()
	step := len(sli) / 4

	sortSli := make([]int, 0)

	c := make(chan []int)
	var wg sync.WaitGroup


	if len(sli) <= 4 {
		wg.Add(1)
		go bubbleSort(sli, c, &wg)
		t := <-c
		sortSli = append(sortSli, t...)
		wg.Wait()
		fmt.Println(sortSli)
		return
	}

	wg.Add(4)
	for i := 0; i < 4; i++ {
		if i > 2 {
			go bubbleSort(sli[i*step:], c, &wg)
		} else {
			go bubbleSort(sli[i*step:i*step+step], c, &wg)
		}

		t := <-c
		sortSli = append(sortSli, t...)

	}
	wg.Wait()
	fmt.Println(sortSli)
}

func readInt() []int {
	sli := make([]int, 0)
	fmt.Println("When you finish press X.")

	for {
		fmt.Printf("Enter the next number: ")

		var str string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			str = scanner.Text()
		}
		if strings.ToLower(str) == "x" {
			break
		}

		n, _ := strconv.Atoi(str)
		sli = append(sli, n)
	}

	return sli
}

func bubbleSort(items []int, c chan []int, wg *sync.WaitGroup){
	fmt.Println("Sort of: ", items)

    for i := 0; i < len(items)-1; i++ {
		for j := 0; j < len(items)-i-1; j++ {
			if items[j] > items[j+1] {
				swap(items, j)
			}
		}
    }

    c <- items
	wg.Done()
}

func swap(items []int, i int) {
	items[i], items[i+1] = items[i+1], items[i]
}
