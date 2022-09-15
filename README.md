# quiz_game
creating a quiz game (learning go by doing projects)
challenges and learnings:
  - concurrency to allow user input in a non-blocking way
  - context to shut down functions mid-loop to make the time-limit work
  - enable customization options by the user via flags

Quiz-game parses a csv in the format: row:"question","solution" and asks the questions, takes in user input and evaluates the answer.
The game closes after a time limit (30s by default) and shows the amount of questions asked and the amount of points that got scored.
The user can specify the time-limit via the -time=XXs flag, the quiz via the -quiz=name.csv flag and the order of questions asked can
be randomized by using the -random flag.

