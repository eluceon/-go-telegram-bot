package main

import (
	"github.com/eluceon/quizwithmebot/internal/app/telegram/app"
	"github.com/eluceon/quizwithmebot/internal/app/telegram/config"
	"log"
	"os"
)

func main() {
	b, err := os.ReadFile("./configs/telegram_config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.ParseConfig(b)
	if err != nil {
		log.Fatal(err)
	}

	tg := app.New(cfg.ApiKeys.Telegram)
	tg.Start()
}
