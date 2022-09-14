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

func Asker(questions, solutions []string, ctx context.Context) (points, amount int) {

	var answer string

	for i := 0; i < len(questions)-1; i++ {
		fmt.Println("What is ", questions[i])
		amount += 1
		fmt.Scanf("%s", &answer)
		if answer == string(solutions[i]) {
			points += 1
		}
	}
	return points, amount
}

func main() {

	var timer time.Duration = 2
	ctx := context.Background()
	ctx1, _ := context.WithTimeout(ctx, timer*time.Second)

	//Seems to work
	questions, solutions := ParseCsv("default.csv")

	points, amount := Asker(questions, solutions, ctx1)
	fmt.Println(points, amount)
}
