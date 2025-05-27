package postgres

import (
	"context"
	"fmt"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Storage struct {
	db *pgxpool.Pool
}

func NewStorage(cfg *config.Config) *Storage {

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		cfg.DB.Host,
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Name,
		cfg.DB.Port)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		panic(err)
	}

	return &Storage{db: pool}
}
