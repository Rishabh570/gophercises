package main

import (
	"flag"
	"fmt"
	"os"
	"quiz-game/utils"
	"time"
)

func main() {
	fileName := flag.String("filename", "check.csv", "name of the csv file")
	timeLimit := flag.Int64("limit", 5, "time limit to solve a problem")
	flag.Parse()

	filePath := fmt.Sprintf("./data/%s", *fileName)

	records := utils.ReadCSV(filePath)
	problems := utils.ParseRecords(records)

	correctAnswers := 0

	fmt.Println("Shall we start the quiz:")
	var startQuiz bool
	fmt.Scanf("%t\n", &startQuiz)
	if !startQuiz {
		fmt.Println("Exiting the quiz")
		os.Exit(0)
	}

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.Question)

		// looks similar to event emitter where we can produce/consume events
		answerCh := make(chan string)

		// Self-executing fn like IIFE in JS, prefix "go" to indicate goroutine
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d.\n", correctAnswers, len(problems))
			return
		case answer := <-answerCh:
			if answer == problem.Answer {
				correctAnswers++
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correctAnswers, len(problems))
}
