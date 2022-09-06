package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func openAndParseCsv(location string) []string {

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

func countPoints(quiz []string) {

	//Compares user input via the console to the answer and counts
	//how many questions were asked, how many were answered correctly.
	//The array of strings that's handed over should have the format ["question", "answer"]
	//and can contain as many question:answer pairs as needed.
	var answer string
	var points int
	var questions int

	for i := 0; i < len(quiz)-1; i += 2 {
		fmt.Println("What is ", quiz[i])
		fmt.Scanf("%s", &answer)

		if answer == string(quiz[i+1]) {
			points += 1
		}
		questions += 1

	}
	fmt.Printf("You scored %d points.\nThere were %d questions in total.\n", points, questions)
}

func main() {
	countPoints(openAndParseCsv("default.csv"))
}
