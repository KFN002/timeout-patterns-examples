package main

import (
	"time"
)

func QuizRunner(questions, answers []string, answerCh chan string) int {
	correctAnswers := 0

	for i, _ := range questions {
		select {
		case answer := <-answerCh:
			if answer == answers[i] {
				correctAnswers++
			}
		case <-time.After(1 * time.Second):
			continue
		}
	}

	return correctAnswers
}
