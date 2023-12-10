package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	getProblems, getError := os.Open("problems.csv")

	if getError != nil {
		fmt.Println("An error occurred while opening CSV file")
	}

	fmt.Println("Successfully opened the CSV file")
	defer getProblems.Close()

	// read CSV file
	fileReader := csv.NewReader(getProblems)
	records, getError := fileReader.ReadAll()
	if getError != nil {
		fmt.Println(getError)
	}
	fmt.Println(records)
}
