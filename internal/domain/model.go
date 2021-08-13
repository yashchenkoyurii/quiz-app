package domain

import (
	"fmt"
	"math/rand"
	"time"
)

type Quiz struct {
	problems []*Problem
}

func (q *Quiz) Problems() []*Problem {
	return q.problems
}

func NewQuiz(problems []*Problem) *Quiz {
	return &Quiz{problems: problems}
}

func (q *Quiz) GetScore() string {
	count := len(q.problems)
	correct := 0

	for _, p := range q.problems {
		if p.isCorrect {
			correct++
		}
	}

	return fmt.Sprintf("Your score: %d/%d", correct, count)
}

func (q *Quiz) Start(problems chan *Problem) {
	defer close(problems)

	for _, p := range q.problems {
		problems <- p
	}
}

func (q *Quiz) Shuffle() {
	problems := &q.problems

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*problems), func(i, j int) {
		(*problems)[i], (*problems)[j] = (*problems)[j], (*problems)[i]
	})
}

type Problem struct {
	question  string
	answer    string
	isCorrect bool
}

func (p *Problem) Question() string {
	return p.question
}

func (p *Problem) IsCorrect() bool {
	return p.isCorrect
}

func NewProblem(question string, answer string) *Problem {
	return &Problem{question: question, answer: answer}
}

func (p *Problem) CheckAnswer(answer string) {
	p.isCorrect = p.answer == answer
}
