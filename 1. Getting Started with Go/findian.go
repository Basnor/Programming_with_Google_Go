package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Enter a string:")

	var str string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		str = scanner.Text()
	}

	lowerStr := strings.ToLower(str)

	if strings.HasPrefix(lowerStr, "i") && strings.Contains(lowerStr, "a") && strings.HasSuffix(lowerStr, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}

	
}