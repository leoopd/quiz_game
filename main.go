package main

import (
	"context"
	"flag"
	"fmt"
	"time"
)

// Quiz-game parses a csv in the format: row:"question","solution" and asks the questions, takes in user input and evaluates the answer.
// The game closes after a time limit (30s by default) and shows the amount of questions asked and the amount of points that got scored.
// The user can specify the time-limit via the -time=XXs flag, the quiz via the -quiz=name.csv flag and the order of questions asked can
// be randomized by using the -random flag.
func main() {

	//Flags, one for the length of the quiz, one for the link to a quiy and one to randomiye the questions.
	definedTime := flag.Int("time", 1, "specifies the time limit for the quiz. Format: 1s/1m/1h")
	quizLink := flag.String("quiz", "default.csv", "specifies the link to the quiz. Format: dir/name.csv")
	shuffleQuestions := flag.Bool("random", false, "randomizes the order of questions")
	flag.Parse()
	questions, solutions := parseCsv(*quizLink)

	//Staring sequence, asks the user to press a specified button and counts down from 3 after.
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

	//Basis to terminate the quiz right when the specified timelimit is reached.
	ctx := context.Background()
	ticker := time.NewTicker(time.Duration(*definedTime) * time.Second)

	points, amount := asker(ctx, ticker, questions, solutions, *shuffleQuestions)

	//Final output that shows the amount of points the user scored and how many questions got asked.
	fmt.Printf("\nYou scored %d out of %d possible points\n", points, amount)
}
