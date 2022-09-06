package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
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

func timer(c chan *time.Time) {
	timer := time.NewTimer(100 * time.Second)
	<-timer.C
}

func countPoints2(quiz []string, pnts chan int, qnts chan int) {

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
			pnts <- points
		}
		questions += 1

		qnts <- questions
	}
}

func main() {

	pnts := make(chan int, 13)
	qnts := make(chan int, 13)
	tmr := make(chan *time.Time)

	go func() {
		countPoints2(openAndParseCsv2("default.csv"), pnts, qnts)
	}()

	go func() {
		timer(tmr)
	}()

	select {
	case <-pnts:
		points := <-pnts
		questions := <-qnts
		fmt.Printf("You scored %d points.\nThere were %d questions in total.\n", points, questions)

	case <-tmr:
		points := <-pnts
		questions := <-qnts
		fmt.Printf("You scored %d points.\nThere were %d questions in total.\n", points, questions)
	}

}
