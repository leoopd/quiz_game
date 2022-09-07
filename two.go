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

func openAndParseCsv2(location string) []string {

	//Opens the .csv file that's specified as the argument
	//and returns the content as an array of strings.
	quizFile, err := os.Open(location)
	if err != nil {
		log.Fatal(err)
	}
	defer quizFile.Close()

	r := csv.NewReader(quizFile)

	var quiz []string

	for {
		row, err := r.Read()
		quiz = append(quiz, row...)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	return quiz
}

func main() {

	var points int
	var questions int
	quiz := openAndParseCsv2("default.csv")
	resultPoints := make(chan int)
	resultQuestions := make(chan int)
	timer := time.NewTicker(100 * time.Second)

	go func() {
		for i := 0; ; {
			select {
			case <-timer.C:
				resultPoints <- points
				resultQuestions <- questions
				close(resultPoints)
				close(resultQuestions)
				return

			default:
				if i > len(quiz)-1 {
					resultPoints <- points
					resultQuestions <- questions
					close(resultPoints)
					close(resultQuestions)
					return
				}

				ch := make(chan string)
				fmt.Println("What is ", quiz[i])

				go func(ch chan string) {
					reader := bufio.NewReader(os.Stdin)
					for {
						s, err := reader.ReadString('\n')
						if err != nil {
							close(ch)
							return
						}
						ch <- s
					}
				}(ch)

			stdinloop:
				for {
					select {
					case answer, ok := <-ch:
						if !ok {
							break stdinloop
						} else {
							if strings.Trim(answer, "\n") == string(quiz[i+1]) {
								points += 1
							}
							questions += 1
							i += 2
						}
					case <-timer.C:
						resultPoints <- points
						resultQuestions <- questions
						close(resultPoints)
						close(resultQuestions)
						return
					}
				}
			}
		}
	}()

	points = <-resultPoints
	questions = <-resultQuestions
	fmt.Printf("You scored %d points.\nThere were %d questions in total.\n", points, questions)
}
