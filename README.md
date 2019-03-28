# Quiz Game

Solution to [jon calhouns gophercise exercise 1](https://github.com/gophercises/quiz)

Program reads quiz questions and answers from CSV file and will then give the quiz to a user, keeping track of how many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards. The quiz is timed and is set to 30 seconds by default. The quiz will stop once the time is up, regardless of whether you had finished the quiz or not

The CSV file is defaulted to problems.csv, but the user is able to customize the filename via a flag.

At the end of the quiz the program outputs the total number of questions correct and how many questions there were in total. Questions given invalid answers are considered incorrect.

The CSV file is in a format like below, where the first column is a question and the second column in the same row is the answer to that question.

```
5+8,13
7+3,10
1+1,2
Capital of france,paris
Capital of germany,berlin
Capital of belguim,brussels
```

At the end of the quiz the program outputs the total number of questions correct and how many questions there were in total. Questions given invalid answers are considered incorrect.

## Futher implementation
1. Add an option (a new flag) to shuffle the quiz order each time it is run.