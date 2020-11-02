package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	option_a string
	option_b string
	option_c string
	option_d string
	answer   string
}

func main() {
	csvFileName := flag.String("csv", "quiz_input.csv", "a csv file in the format of 'Questions,Four Options,Answers'")
	timeLimit := flag.Int("limit", 20, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("Failed to parse the provided CSV file")
	}
	problems := parseLines(lines)
	// fmt.Println(problems)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Q%d - %s \n", i+1, p.question)
		fmt.Printf("%s\n%s\n%s\n%s\n", p.option_a, p.option_b, p.option_c, p.option_d)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Printf("Answer: ")
			fmt.Scanf("%s\n", &answer)
			fmt.Printf("--------------\n")
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
			return
		case answer := <-answerCh:
			answer = strings.ToLower(answer)
			if answer == p.answer {
				correct++
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			option_a: strings.TrimSpace(line[1]),
			option_b: strings.TrimSpace(line[2]),
			option_c: strings.TrimSpace(line[3]),
			option_d: strings.TrimSpace(line[4]),
			answer:   strings.TrimSpace(line[5]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
