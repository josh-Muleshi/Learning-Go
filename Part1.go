package main

import (
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
}
