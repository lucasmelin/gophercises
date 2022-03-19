// Read in math problems from a CSV, and prompt for
// solutions from command line.
// flags, csv, os packages
// channels, coroutines and the time package for the timer
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	filenamePtr := flag.String("filename", "addition_problems.csv", "a CSV file with the format `question,answer`")
	flag.Parse()

	file, err := os.Open(*filenamePtr)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *filenamePtr))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse the CSV file: %s", *filenamePtr))
	}
	problems := parseLines(lines)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		var answer string
		// Scanf will trim any whitespace from the answer
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	// We know ahead of time the size of the slice
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
