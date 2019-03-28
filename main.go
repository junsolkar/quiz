package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// DEFAULTTIMELIMIT sets the default time limit for the quiz
const DEFAULTTIMELIMIT = 30 //seconds

type problem struct {
	question string
	answer   string
}

func main() {
	fileName := flag.String("csv", "problems.csv", "file with math questions in format 'question,answer'")
	quizTimeLimit := flag.Int("limit", DEFAULTTIMELIMIT, "the time limit for the quiz")

	flag.Parse()
	file, err := os.Open(*fileName)
	if err != nil {
		exit(fmt.Sprintf("failed to open csv file: %s\n", *fileName))
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse csv file.")
	}

	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*quizTimeLimit) * time.Second)
	defer timer.Stop()

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problems #%d: %s = ", i+1, p.question)

		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nIn total you scored %d out of %d \n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.answer {
				correct++
			}
		}
	}

	fmt.Printf("In total you scored %d out of %d \n", correct, len(problems))
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

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}