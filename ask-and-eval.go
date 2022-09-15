package main

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Asks the question, takes in the user Input in a non-blocking way and evaluates if the answer is correct.
//The function returns two counters, points for every correct answer and amount for every question asked.
func asker(ctx context.Context, ticker *time.Ticker, questions, solutions []string, shuffleQuestions bool) (points, amount int) {

	ctx, cancel := context.WithCancel(ctx)

	ch := make(chan int)
	amount = 1

	//Mechanism to randomize the order of the questions asked if the user wants to.
	var order []int
	if shuffleQuestions {
		for i := range questions {
			order = append(order, i)
		}
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(order), func(i, j int) { order[i], order[j] = order[j], order[i] })
	} else {
		for i := range questions {
			order = append(order, i)
		}
	}

	//Non-blocking way of taking in user input by launching a GO-Routine and sending the input via a channel back to the function. 
	go func(ctx context.Context, ch chan int, order []int) {

		var tmp int
		reader := bufio.NewReader(os.Stdin)
		for i := 0; i < len(questions); i++ {
			fmt.Printf("What is: %s?\n", questions[order[i]])
			s, err := reader.ReadString('\n')
			if err != nil {
				close(ch)
				return
			} else if strings.Trim(s, "\n") == solutions[order[i]] {
				tmp = 1
			} else {
				tmp = 0
			}
			ch <- tmp
		}
	}(ctx, ch, order)

//Takes in the user input that was sent through the channel and evaluates if the answer is correct.
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
			cancel()
			return points, amount
		}
	}
	cancel()
	return points, amount
}
