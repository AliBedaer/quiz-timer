package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var correctAnswer = 0

func main() {
	csvFilename := flag.String("csv", "problems.csv", "the csv file that contain quizes")
	timeLimit := flag.Int("limit", 30, "the time needed to solve this problems in second")
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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, quiz := range problems {

		getanswerChan := make(chan string)
		go getAnswer(quiz.question, i, getanswerChan)

		select {
		case <-timer.C:
			fmt.Printf("You Answerd %d out of %d \n", correctAnswer, len(problems))
			return
		case answer := <-getanswerChan:
			checkAnswer(quiz.answer, answer)
		}
	}

	fmt.Printf("You Answerd %d out of %d \n", correctAnswer, len(problems))

}

// # get answer function

func getAnswer(question string, index int, getAnswer chan string) {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Printf("Problem %d: %s = ", index+1, question)
	reader.Scan()
	getAnswer <- reader.Text()
}

// # checkQuiz function

func checkAnswer(answer string, response string) {
	if strings.TrimSpace(response) == answer {
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
