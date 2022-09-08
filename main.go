package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// Gets a location for a csv file that contains the questions and answers in the format: question,solutions
// and 2 channels from main, q to send the questions to ask.go and s to send the solutions to eval.go.
func ParseCsv(location string, s chan []string) (questions []string) {

	//Opens the .csv file that's specified as the argument and sends
	//the content as arrays of strings, split into questions and answers.
	quizFile, err := os.Open(location)
	if err != nil {
		log.Fatal(err)
	}
	defer quizFile.Close()

	r := csv.NewReader(quizFile)
	var solutions []string

	// var questions []string
	// var solutions []string

	for {
		row, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		questions = append(questions, row[0])
		solutions = append(solutions, row[1])

	}
	s <- solutions
	return questions
}

// Gets the questions via the channel q from parse.go and 2 channels from main, answers to send the answers given by the user to eval.go
// and questionCount to send the amount of questions asked to main.go.
func Asker(questions []string, q chan int) (answersGiven []string) {

	ticker := time.NewTicker(time.Second * 5)
	var ch chan string
	var questionsAsked int

	for i := 0; i < len(questions); i++ {

		fmt.Println("What is ", questions[i])
		questionsAsked += 1

		go func(ch chan string, i int) {
			fmt.Println("123")
			reader := bufio.NewReader(os.Stdin)
			for {
				s, err := reader.ReadString('\n')
				if err != nil {
					log.Fatal(err)
					close(ch)
					return
				}
				fmt.Println(s)
				ch <- s
			}
		}(ch, i)

	stdinloop:
		for {
			select {
			case answer, ok := <-ch:
				if !ok {
					break stdinloop
				} else {
					answersGiven = append(answersGiven, strings.Trim(answer, "\n"))
					fmt.Println(answersGiven)
				}
			case <-ticker.C:
				q <- questionsAsked
				return answersGiven
			}
		}
	}
	q <- questionsAsked
	return answersGiven
}

// Gets the solutions sent from parse.go via channel s and the answers from ask.go sent via channel answers.
// Compares the answers to the solutions and returns the amount of points the user scored.
func Eval(answersGiven []string, s chan []string) (points int) {

	solutions := <-s

	for i := 0; i < len(answersGiven); i++ {
		if answersGiven[i] == solutions[i] {
			points += 1
		}
	}

	return points
}

func main() {

	s := make(chan []string)
	q := make(chan int)

	// answers := make(chan []string)
	// questionCount := make(chan int)

	// ParseCsv("default.csv", q, s)
	// Asker(q, answers, questionCount)

	// points := Eval(s, answers)
	// questionsAsked := <-questionCount

	// fmt.Printf("Points scored: %d\nQuestions asked: %d", points, questionsAsked)

	fmt.Println(Asker(ParseCsv("default.csv", s), q))

	// questions := <-q

	// fmt.Printf("questions: %v, solutions: %v\n", questions, points)

}
