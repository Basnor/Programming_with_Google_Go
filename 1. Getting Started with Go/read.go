package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type Record struct {
	fname, lname string
}

func main() {
	var filename string
	fmt.Print("Enter a filename: ")
	fmt.Scan(&filename)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceOfRecords := make([]Record, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var a Record = splitStringToRecord(scanner.Text())
		sliceOfRecords = append(sliceOfRecords, a)
	}
	f.Close()

	for i, rec := range sliceOfRecords {
		fmt.Printf("%v. %s %s\n", i+1, rec.fname, rec.lname)
	}
}

func splitStringToRecord(s string) Record {
	//fmt.Println(s)
	var rec Record
	rec.fname = strings.Split(s, " ")[0]
	rec.lname = strings.Split(s, " ")[1]
	return rec
}
