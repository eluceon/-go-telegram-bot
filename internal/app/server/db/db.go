package db

import (
	"context"
	"fmt"
	"github.com/eluceon/quizwithmebot/internal/app/server/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
)

func New(ctx context.Context) (*pgxpool.Pool, error) {
	b, err := os.ReadFile("./configs/server_config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.ParseConfig(b)
	if err != nil {
		log.Fatal(err)
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBName, cfg.DB.Sslmode)
	return pgxpool.Connect(ctx, dsn)
}
