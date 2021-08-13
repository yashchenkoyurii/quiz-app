package infrastructure

import (
	"encoding/csv"
	"github.com/yashchenkoyurii/quiz-app/config"
	"github.com/yashchenkoyurii/quiz-app/internal/domain"
	"log"
	"os"
)

type QuizRepository struct {
	config *config.Config
	csv    *csv.Reader
}

func NewQuizRepository(config *config.Config) *QuizRepository {
	return &QuizRepository{config: config}
}

func (q *QuizRepository) Get() *domain.Quiz {
	file, err := os.Open(q.config.Filename)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	lines, err := q.csv.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	var problems []*domain.Problem

	for _, line := range lines {
		problems = append(problems, domain.NewProblem(line[0], line[1]))
	}

	return domain.NewQuiz(problems)
}
