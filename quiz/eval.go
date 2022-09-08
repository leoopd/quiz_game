package quiz

// Gets the solutions sent from parse.go via channel s and the answers from ask.go sent via channel answers.
// Compares the answers to the solutions and returns the amount of points the user scored.
func Eval(s, answers chan []string) (points int) {

	solutions := <-s
	answersGiven := <-answers

	for i := 0; i < len(answersGiven); i++ {
		if answersGiven[i] == solutions[i] {
			points += 1
		}
	}

	return points
}
