package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")

	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("failed to open the csv file: %s", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll() // not ok for huge files
	if err != nil {
		exit("failed to parse the provided csv file")
	}

	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)

		var answer string
		fmt.Scanf("%s\n", &answer) // trims the input

		if answer == p.a {
			correct++
		}
	}

	fmt.Printf("Yous scored %d out of %d.\n", correct, len(problems))
}

// parseLines does not use append() because we know the length of the slice we want.
// append() resize logic has a performance impact instead
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
