package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

//Parses the csv specified by the location and returns a list of questions and solutions if the format is correct.
func parseCsv(location string) (questions, solutions []string) {

	//Opens the file specified via the location. Can be set by the user by using the flag.
	quizFile, err := os.Open(location)
	if err != nil {
		log.Fatal(err)
	}
	defer quizFile.Close()

	r := csv.NewReader(quizFile)

	//Reads the csv row by row and appends the corresponding parts to the questions/solutions slices.
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
