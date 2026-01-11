package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"judo_stats_site/internal/app"
	"judo_stats_site/internal/config"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	cfg := config.MustLoad(logger)

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
