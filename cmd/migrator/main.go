package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Nikita-Mihailuk/smartwayTestTask/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var migrationsPath, migrationsTable string

	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations directory")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name for migrations table")
	flag.Parse()

	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	cfg := config.GetConfig()
	dbURL := fmt.Sprintf("%s:%s@%s:%s/%s",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("postgres://%s?x-migrations-table=%s&sslmode=disable", dbURL, migrationsTable))

	if err != nil {
		panic(fmt.Errorf("failed to initialize migrator: %w", err))
	}

	if err = m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")
			return
		}
		panic(fmt.Errorf("migration failed: %w", err))
	}

	fmt.Println("migrations applied successfully")
}
