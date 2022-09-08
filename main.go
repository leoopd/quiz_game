package main

import (
	"fmt"
)

func main() {

	q := make(chan []string)
	s := make(chan []string)

	answers := make(chan []string)
	questionCount := make(chan int)

	quiz.ParseCsv("default.csv", q, s)
	quiz.Asker(q, answers, questionCount)

	points := quiz.Eval(s, answers)
	questionsAsked := <-questionCount

	fmt.Printf("Points scored: %d\nQuestions asked: %d", points, questionsAsked)
}
