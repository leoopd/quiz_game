package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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

	ctx, cancel := context.WithCancel(ctx)

	ch := make(chan int)
	amount = 1

	// for i := 0; i < len(questions); i++ {
	go func(ctx context.Context, ch chan int) {

		var tmp int
		reader := bufio.NewReader(os.Stdin)
		for i := 0; i < len(questions); i++ {
			fmt.Println("What is: ", questions[i])
			s, err := reader.ReadString('\n')
			if err != nil {
				close(ch)
				return
			} else if strings.Trim(s, "\n") == solutions[i] {
				tmp = 1
			} else {
				tmp = 0
			}
			ch <- tmp
		}
	}(ctx, ch)
	// }

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
			return points, amount
			cancel()
		}
	}
	return points, amount
}

func main() {

	var definedTime time.Duration = 2
	ctx := context.Background()
	ticker := time.NewTicker(definedTime * time.Second)

	//Seems to work
	questions, solutions := ParseCsv("default.csv")

	points, amount := Asker(ctx, ticker, questions, solutions)
	fmt.Printf("\nYou scored %d out of %d possible points\n", points, amount)
}
