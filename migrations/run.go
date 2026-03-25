package migrations

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func Run(connStr string) error {
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return fmt.Errorf("открытие соединения для миграций: %w", err)
	}
	defer db.Close()

	goose.SetBaseFS(Files)

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("настройка диалекта goose: %w", err)
	}

	if err := goose.Up(db, "."); err != nil {
		return fmt.Errorf("выполнение миграций: %w", err)
	}

	return nil
}
