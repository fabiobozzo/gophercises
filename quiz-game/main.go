package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	shuffle := flag.Bool("shuffle", false, "if set the problems are shown in random order")
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
	if *shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(problems), func(i, j int) {
			problems[i], problems[j] = problems[j], problems[i]
		})
	}

	timer := time.NewTicker(time.Duration(*timeLimit) * time.Second)
	correct := 0

problemLoop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)

		answerCH := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer) // trims the input
			answerCH <- answer
		}()

		select {
		case <-timer.C:
			break problemLoop
		case answer := <-answerCH:
			if answer == p.a {
				correct++
			}
		}
	}

	fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
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
