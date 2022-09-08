package quiz

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Gets the questions via the channel q from parse.go and 2 channels from main, answers to send the answers given by the user to eval.go
// and questionCount to send the amount of questions asked to main.go.
func Asker(q, answers chan []string, questionCount chan int) {

	ticker := time.NewTicker(time.Second * 5)
	var questionsAsked int
	var ch chan string
	var answersGiven []string
	questions := <-q

	for i := 0; i < len(questions); i++ {
		go func(ch chan string, i int) {
			fmt.Println("What is ", questions[i])
			questionsAsked += 1

			reader := bufio.NewReader(os.Stdin)
			for {
				s, err := reader.ReadString('\n')
				if err != nil {
					close(ch)
					return
				}
				ch <- s
			}
		}(ch, i)

	stdinloop:
		for {
			select {
			case answer, ok := <-ch:
				if !ok {
					answers <- answersGiven
					questionCount <- questionsAsked
					break stdinloop
				} else {
					answersGiven = append(answersGiven, strings.Trim(answer, "\n"))
				}
			case <-ticker.C:
				answers <- answersGiven
				questionCount <- questionsAsked
				return
			}
		}
	}
}
