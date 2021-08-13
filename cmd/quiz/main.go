package main

import (
	"flag"
	"github.com/yashchenkoyurii/quiz-app/config"
	"github.com/yashchenkoyurii/quiz-app/internal/application"
	infrastructure "github.com/yashchenkoyurii/quiz-app/internal/infrastructure/repository"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		filename string
		limit    int
		shuffle  bool
	)

	flag.StringVar(&filename, "filename", "quiz.csv", "Path to csv file")
	flag.IntVar(&limit, "limit", 300, "Duration of the quiz")
	flag.BoolVar(&shuffle, "shuffle", false, "Shuffle questions")
	flag.Parse()

	quit := make(chan os.Signal)
	conf := config.NewConfig(filename, limit, shuffle)
	quizRepository := infrastructure.NewQuizRepository(conf)
	quizService := application.NewQuizService(quizRepository, conf)

	go quizService.StartQuiz(quit)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
}
