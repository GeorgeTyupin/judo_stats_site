package app

import (
	"context"
	"fmt"
	"judo_stats_site/internal/config"
	"judo_stats_site/internal/handlers"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
)

type ServerApp struct {
	AppName string
	logger  *slog.Logger
	server  *http.Server
	cfg     *config.Config
}

func NewApp(log *slog.Logger, cfg *config.Config) *ServerApp {
	appName := "HTTPServer"
	log = log.With(slog.String("app", appName))

	handler := registerHandlers()

	server := &http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: handler,
	}

	app := &ServerApp{
		AppName: appName,
		logger:  log,
		server:  server,
		cfg:     cfg,
	}

	return app
}

func (app *ServerApp) Run() error {
	const op = "app.Run"
	logger := app.logger.With(slog.String("op", op))
	logger.Info("Запуск сервера...", slog.String("address", app.server.Addr))

	return app.server.ListenAndServe()
}

func (app *ServerApp) Shutdown() {
	const op = "app.Shutdown"
	logger := app.logger.With(slog.String("op", op))
	logger.Info("Остановка сервера...")

	ctx, cancel := context.WithTimeout(context.Background(), app.cfg.HTTPServer.Timeouts.Shutdown)
	defer cancel()

	if err := app.server.Shutdown(ctx); err != nil {
		warning := fmt.Sprintf("Сервер не успел завершиться за %v. Завершаем принудительно.", app.cfg.HTTPServer.Timeouts.Shutdown)
		logger.Warn(warning, slog.String("error", err.Error()))

		if err := app.server.Close(); err != nil {
			closeErr := fmt.Sprintf("Ошибка критическая остановки сервера: %v", err)
			logger.Error(closeErr, slog.String("error", err.Error()))
		}
	}
}

func registerHandlers() *chi.Mux {
	r := chi.NewRouter()

	// Статические файлы
	fileServer := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Страницы
	r.Get("/", handlers.Index)
	r.Get("/judoka", handlers.Judoka)
	r.Get("/tournament", handlers.Tournament)
	r.Get("/sportclub", handlers.SportClub)
	r.Get("/city", handlers.City)
	r.Get("/contribute", handlers.Contribute)

	// TODO Реализовать API endpoints для форм
	// r.Post("/api/contribute/tournament", handlers.ContributeTournament)
	// r.Post("/api/contribute/athlete", handlers.ContributeAthlete)
	// r.Post("/api/contribute/error", handlers.ContributeError)

	return r
}
