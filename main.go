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

// DEFAULTCSVFILE sets the default csv file for the program to read from
const DEFAULTCSVFILE = "problems.csv"

type problem struct {
	question string
	answer   string
}

func main() {
	fileName := flag.String("csv", DEFAULTCSVFILE, "file with questions in format 'question,answer'")
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

	correct := AnswerProblem(problems, timer)

	fmt.Printf("\nIn total you scored %d out of %d \n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   RemoveSpaceAndLowerCase(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// RemoveSpaceAndLowerCase removes white space and lowercases the input string
func RemoveSpaceAndLowerCase(s string) string {
	return strings.TrimSpace(strings.ToLower(s))
}

// AnswerProblem will read through the array of problems and calculate how many were correct
func AnswerProblem(problemList []problem, timer *time.Timer) int {
	correct := 0
	for i, p := range problemList {
		fmt.Printf("Problem #%d: %s = ", i+1, p.question)

		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			return correct
		case answer := <-answerCh:
			if RemoveSpaceAndLowerCase(answer) == RemoveSpaceAndLowerCase(p.answer) {
				correct++
			}
		}
	}

	return correct
}
