package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
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

// func main() {

// 	var points int
// 	var questions int
// 	var answers []string
// 	quiz := openAndParseCsv2("default.csv")
// 	ans := make(chan []string)
// 	timer := time.NewTicker(5 * time.Second)

// 	go func() {
// 		for i := 0; ; {
// 			ch := make(chan string)
// 			fmt.Println("What is ", quiz[i])

// 			go func(ch chan string) {
// 				reader := bufio.NewReader(os.Stdin)
// 				for {
// 					s, err := reader.ReadString('\n')
// 					if err != nil {
// 						close(ch)
// 						return
// 					}
// 					ch <- s
// 				}
// 			}(ch)

// 		stdinloop:
// 			for {
// 				select {
// 				case answer, ok := <-ch:
// 					if !ok {
// 						break stdinloop
// 					} else {
// 						answers = append(answers, strings.Trim(answer, "\n"))
// 					}
// 				case <-timer.C:
// 					ans <- answers
// 					return
// 				}
// 			}
// 		}
// 	}()

// 	answers = <-ans
// 	fmt.Println(answers)

// 	for i := 1; i < len(answers); i += 2 {
// 		if answers[i] == quiz[i*2] {
// 			points += 1
// 		}
// 		questions += 1
// 	}
// 	fmt.Printf("You scored %d points.\nThere were %d questions in total.\n", points, questions)
// }
