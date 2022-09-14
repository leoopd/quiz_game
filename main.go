package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	questions, solutions := parseCsv("default.csv")

	fmt.Println("Please press 'enter' to start the quiz")
	fmt.Scanln()
	fmt.Println("The quiz starts in 3...")
	time.Sleep(1 * time.Second)
	fmt.Println("2...")
	time.Sleep(1 * time.Second)
	fmt.Println("1...")
	time.Sleep(1 * time.Second)
	fmt.Println("GO!")
	time.Sleep(333 * time.Millisecond)

	var definedTime time.Duration = 2
	ctx := context.Background()
	ticker := time.NewTicker(definedTime * time.Second)

	points, amount := asker(ctx, ticker, questions, solutions)

	fmt.Printf("\nYou scored %d out of %d possible points\n", points, amount)
}
