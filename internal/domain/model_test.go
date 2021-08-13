package domain_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yashchenkoyurii/quiz-app/internal/domain"
	"math/rand"
	"testing"
)

func TestNewProblem(t *testing.T) {
	problems := []*domain.Problem{
		domain.NewProblem("2+2", "4"),
		domain.NewProblem("4+4", "8"),
	}

	for _, problem := range problems {
		assert.IsType(t, &domain.Problem{}, problem)
	}
}

func TestProblem_CheckAnswer(t *testing.T) {
	testCases := []struct {
		problem   *domain.Problem
		answer    string
		isCorrect bool
	}{
		{domain.NewProblem("2+2", "4"), "4", true},
		{domain.NewProblem("2+3", "5"), "4", false},
		{domain.NewProblem("2+3", "5"), "3", false},
	}

	for _, tc := range testCases {
		tc.problem.CheckAnswer(tc.answer)
		assert.Equal(t, tc.problem.IsCorrect(), tc.isCorrect)
	}
}

func TestQuiz_GetScore(t *testing.T) {
	problems := []*domain.Problem{
		domain.NewProblem("2+2", "4"),
		domain.NewProblem("2+3", "5"),
		domain.NewProblem("2+3", "5"),
	}

	quiz := domain.NewQuiz(problems)

	assert.Equal(t, quiz.GetScore(), fmt.Sprintf("Your score: 0/%d", len(problems)))
}

func TestQuiz_Start(t *testing.T) {
	problems := []*domain.Problem{
		domain.NewProblem("2+2", "4"),
		domain.NewProblem("2+3", "5"),
		domain.NewProblem("2+3", "5"),
	}

	quiz := domain.NewQuiz(problems)
	ch := make(chan *domain.Problem)
	count := 0
	go quiz.Start(ch)

	for p := range ch {
		assert.IsType(t, &domain.Problem{}, p)
		count++
	}

	assert.Equal(t, len(problems), count)
}

func TestQuiz_Shuffle(t *testing.T) {
	problems := []*domain.Problem{
		domain.NewProblem("2+2", "4"),
		domain.NewProblem("2+6", "8"),
		domain.NewProblem("2+3", "5"),
	}
	quiz := domain.NewQuiz(problems)

	rand.Seed(42)
	p1 := quiz.Problems()[0]
	quiz.Shuffle()
	p2 := quiz.Problems()[0]

	assert.NotEqual(t, p1, p2)
}
