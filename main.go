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
func ParseCsv(location string) (questions, solutions []string) {

	//Opens the .csv file that's specified as the argument and sends
	//the content as arrays of strings, split into questions and answers.
	quizFile, err := os.Open(location)
	if err != nil {
		log.Fatal(err)
	}
	defer quizFile.Close()

	r := csv.NewReader(quizFile)

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
	return questions, solutions
}

// Gets the questions via the channel q from parse.go and 2 channels from main, answers to send the answers given by the user to eval.go
// and questionCount to send the amount of questions asked to main.go.
func Asker(questions []string) (answersGiven []string, questionsAsked int) {

	ticker := time.NewTicker(time.Second * 5)
	var ch chan string

	// for i := 0; i < len(questions); i++ {

	go func(ch chan string) {
		reader := bufio.NewReader(os.Stdin)
		for i := 0; i < len(questions); i++ {

			fmt.Println("What is ", questions[i])
			questionsAsked += 1

			s, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
				close(ch)
				return
			}
			ch <- s
		}
	}(ch)

	for {
		select {

		case answer, ok := <-ch:
			fmt.Println(ok)
			if !ok {
				return answersGiven, questionsAsked
			} else {
				fmt.Println("hier2")

				answersGiven = append(answersGiven, strings.Trim(answer, "\n"))
			}
		case <-ticker.C:
			return answersGiven, questionsAsked
		}
	}
	fmt.Println("hier")
	return answersGiven, questionsAsked
}

// }

// Gets the solutions sent from parse.go via channel s and the answers from ask.go sent via channel answers.
// Compares the answers to the solutions and returns the amount of points the user scored.
func Eval(answersGiven, solutions []string) (points int) {

	for i := 0; i < len(answersGiven); i++ {
		if answersGiven[i] == solutions[i] {
			points += 1
		}
	}

	return points
}

func main() {

	//Seems to work
	questions, solutions := ParseCsv("default.csv")

	//
	answers, questionsAsked := Asker(questions)
	fmt.Println(answers, questionsAsked)

	points := Eval(answers, solutions)
	fmt.Println(points)
}
