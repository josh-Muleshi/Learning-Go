package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Define flags for the CSV file and timer
	csvFileName := flag.String("csv", "problems.csv", "a CSV file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	// Open the CSV file
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}
	defer file.Close()

	// Parse the CSV file
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	// Process the quiz
	problems := parseLines(lines)
	correct := 0

	fmt.Println("Press Enter to start the quiz.")
	fmt.Scanln()

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

problemLoop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTime's up!")
			break problemLoop
		case answer := <-answerCh:
			if answer == p.answer {
				correct++
			}
		}
	}

	fmt.Printf("You got %d out of %d questions correct.\n", correct, len(problems))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

type problem struct {
	question string
	answer   string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}
