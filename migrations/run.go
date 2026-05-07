package migrations

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func Run(connStr string, logger *slog.Logger) error {
	const op = "migrations.Run"
	logger = logger.With(slog.String("op", op))

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return fmt.Errorf("открытие соединения для миграций: %w", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error("не удалось закрыть соединение для миграций", slog.String("error", err.Error()))
		}
	}()

	provider, err := goose.NewProvider(goose.DialectPostgres, db, Files,
		goose.WithTableName("public.goose_db_version"),
	)
	if err != nil {
		return fmt.Errorf("создание goose provider: %w", err)
	}

	if _, err := provider.Up(context.Background()); err != nil {
		return fmt.Errorf("выполнение миграций: %w", err)
	}

	return nil
}
