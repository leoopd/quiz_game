package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func ParseCsv(location string) (questions, solutions []string) {

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

func Asker(ctx context.Context, ticker *time.Ticker, questions, solutions []string) (points, amount int) {

	ch := make(chan int)
	amount = 1

	for i := 0; i < len(questions); i++ {
		go func(ctx context.Context, ch chan int, i int) {

			reader := bufio.NewReader(os.Stdin)
			for {
				fmt.Println("What is: ", questions[i])
				s, err := reader.ReadString('\n')
				if err != nil {
					close(ch)
					return
				} else if s == solutions[i] {
					ch <- 1
				} else {
					ch <- 0
				}
			}
		}(ctx, ch, i)
	}

stdinloop:
	for {
		select {
		case eval, ok := <-ch:
			if !ok {
				break stdinloop
			} else {
				if eval == 1 {
					points += 1
					amount += 1
				} else {
					amount += 1
				}
			}
		case <-ticker.C:

		}
	}
	fmt.Println("Done, stdin must be closed")
}

func main() {

	var definedTime time.Duration
	ctx := context.Background()
	ctx, _ = context.WithCancel(ctx)
	ticker := time.NewTicker(definedTime * time.Second)

	//Seems to work
	questions, solutions := ParseCsv("default.csv")

	points, amount := Asker(ctx, ticker, questions, solutions)
	fmt.Println(points, amount)
}
