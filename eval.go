package main

// Gets the solutions sent from parse.go via channel s and the answers from ask.go sent via channel answers.
// Compares the answers to the solutions and returns the amount of points the user scored.
func Eval(solutions, answers []string) (points int) {


	for i := 0; i < len(answers); i++ {
		if answers[i] == solutions[i] {
			points += 1
		}
	}

	return points
}
