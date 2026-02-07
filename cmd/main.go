package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"judo_stats_site/internal/app"
	"judo_stats_site/internal/config"
	"judo_stats_site/internal/repository"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	cfg := config.MustLoad(logger)

	repoCtx, repoCtxCancelFunc := context.WithTimeout(context.Background(), cfg.Database.DBIdle)
	defer repoCtxCancelFunc()

	pgRepo, err := repository.NewDBRepository(repoCtx, &cfg.Database, logger)
	if err != nil {
		logger.Error("Ошибка подключения к базе данных", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer pgRepo.Close()

	application := app.NewApp(logger, cfg)

	signalCh := make(chan os.Signal, 2)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	errorCh := make(chan error, 1)

	go func() {
		if err := application.Run(); err != nil {
			errorCh <- err
		}
	}()

	select {
	case sig := <-signalCh:
		logger.Info("Получен сигнал завершения", slog.String("signal", sig.String()))
	case err := <-errorCh:
		logger.Error("Ошибка запуска сервера", slog.String("error", err.Error()))
	}

	application.Shutdown()
}
