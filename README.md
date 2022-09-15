# quiz_game
creating a quiz game (learning go by doing projects)
challenges and learnings:
  - concurrency to allow user input in a non-blocking way
  - context to shut down functions mid-loop to make the time-limit work
  - enable customization options by the user via flags

Quiz-game parses a csv in the format: row:"question","solution" and asks the questions, takes in user input and evaluates the answer. The game closes after a time limit (30s by default) and shows the amount of questions asked and the amount of points that got scored. <br>
Flags that can be used:
  - time (example: -time=10s), sets the time-limit the user has to answer questions
  - quiz (example: -quiz=quizzes/my_quiz.csv), enables the user to feed the game a custom quiz to play with
  - random (example: -random), shuffles the order of the questions asked
