package application

import (
	"fmt"
	"github.com/yashchenkoyurii/quiz-app/config"
	"github.com/yashchenkoyurii/quiz-app/internal/domain"
	"log"
	"os"
	"time"
)

type QuizService struct {
	repository domain.QuizRepository
	config     *config.Config
}

func NewQuizService(repository domain.QuizRepository, conf *config.Config) *QuizService {
	return &QuizService{repository: repository, config: conf}
}

func (s QuizService) StartQuiz(quit chan os.Signal) {
	var input string
	quizChan := make(chan *domain.Problem)
	quiz := s.repository.Get()

	fmt.Println(s.config.Shuffle)
	if s.config.Shuffle {
		quiz.Shuffle()
	}

	timer := time.NewTimer(time.Duration(s.config.Limit) * time.Second)

	go func() {
		<-timer.C

		fmt.Println("Time is out!")
		fmt.Println(quiz.GetScore())
		quit <- os.Interrupt
	}()
	go quiz.Start(quizChan)
	fmt.Println("Quiz started!")

	for problem := range quizChan {
		fmt.Printf("Problem: %s\nYour answer: ", problem.Question())

		if _, err := fmt.Scanln(&input); err != nil {
			log.Println(err)
		}

		problem.CheckAnswer(input)
	}

	fmt.Println(quiz.GetScore())
	quit <- os.Interrupt
}
