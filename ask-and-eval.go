package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"
)

func asker(ctx context.Context, ticker *time.Ticker, questions, solutions []string) (points, amount int) {

	ctx, cancel := context.WithCancel(ctx)

	ch := make(chan int)
	amount = 1

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
