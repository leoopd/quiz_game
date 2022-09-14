package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	var definedTime time.Duration = 2
	ctx := context.Background()
	ticker := time.NewTicker(definedTime * time.Second)

	questions, solutions := ParseCsv("default.csv")

	points, amount := Asker(ctx, ticker, questions, solutions)
	fmt.Printf("\nYou scored %d out of %d possible points\n", points, amount)
}
