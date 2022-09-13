package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
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


func Asker(questions, solutions []string, ctx xontext.Context) (points, amount int){

	for i := 0; i < len(questions); i++ {
		go func(ch chan string, i int) {
			fmt.Println(questions[i])
			amount += 1

			reader := bufio.NewReader(os.Stdin)
			for {
				s, err := reader.ReadString('\n')
				if err != nil {
					close(ch)
					return
				}
				ch <- s
			}
		}(ch, i)

		for {
			select {
			case answer, ok := <-ch:
				if !ok {
					answers <- answersGiven
					questionCount <- questionsAsked
					os.Exit()
				} else {
					answersGiven = append(answersGiven, strings.Trim(answer, "\n"))
				}
			case <-ctx.done:
				answers <- answersGiven
				questionCount <- questionsAsked
				return
			}
		}
	}
}

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

	var timer time.Duration = 2
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, timer*time.Second)

	//Seems to work
	questions, solutions := ParseCsv("default.csv")

	go func(){
		if <- 
	}
 
	points, amount := Asker(questions, solutions, ctx)
	fmt.Println(points, amount)
}
