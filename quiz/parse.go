package quiz

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

// Gets a location for a csv file that contains the questions and answers in the format: question,solutions
// and 2 channels from main, q to send the questions to ask.go and s to send the solutions to eval.go.
func ParseCsv(location string, q, s chan []string) {

	//Opens the .csv file that's specified as the argument and sends
	//the content as arrays of strings, split into questions and answers.
	quizFile, err := os.Open(location)
	if err != nil {
		log.Fatal(err)
	}
	defer quizFile.Close()

	r := csv.NewReader(quizFile)

	var questions []string
	var solutions []string

	for {
		row, err := r.Read()

		questions = append(questions, row[0])
		solutions = append(solutions, row[1])
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	q <- questions
	s <- solutions
}
