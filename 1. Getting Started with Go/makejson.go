package main

import (
	"bufio"
	"fmt"
	"os"
	"encoding/json"
)

func main() {
	nameByAddress := make(map[string]string)

	var name, address string
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Printf("enter the name:")
	if scanner.Scan() {
		name = scanner.Text()
	}

	fmt.Print("enter the address:")
	if scanner.Scan() {
		address = scanner.Text()
	}

	nameByAddress["name"] = name
	nameByAddress["address"] = address
	data, err := json.Marshal(nameByAddress)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print("json:")
	fmt.Println(string(data))
}
