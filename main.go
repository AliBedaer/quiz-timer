package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

var correctAnswer = 0

func main() {
	csvFilename := flag.String("csv", "problems.csv", "the csv file that contain quizes")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the %s file \n", *csvFilename))
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("failed to Parse file")
	}

	problems := parseLines(lines)

	for i, quiz := range problems {
		checkQuiz(quiz.question, quiz.answer, i)
	}

	fmt.Printf("You Answerd %d out of %d \n", correctAnswer, len(problems))
}

// # checkQuiz function

func checkQuiz(question, answer string, index int) {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Printf("Problem %d: %s = \n", index+1, question)
	reader.Scan()

	if strings.TrimSpace(reader.Text()) == answer {
		correctAnswer++
	}

}

// # parseLines

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		qu := problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
		ret[i] = qu
	}
	return ret
}

// #struct

type problem struct {
	question string
	answer   string
}

// # exit function

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
