package app

import (
	"context"
	"fmt"
	"judo_stats_site/internal/config"
	"judo_stats_site/internal/handlers"
	"judo_stats_site/internal/repository"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/crypto/acme/autocert"
)

const component = "HTTPServer"

type ServerApp struct {
	logger       *slog.Logger
	server       *http.Server
	httpRedirect *http.Server
	cfg          *config.Config
	autocertMgr  *autocert.Manager
	db           *repository.DBRepository
}

func NewApp(logger *slog.Logger, cfg *config.Config, repo *repository.DBRepository) *ServerApp {
	logger = logger.With(slog.String("place", component))

	handler := registerHandlers(repo, logger)

	app := &ServerApp{
		logger: logger,
		cfg:    cfg,
		db:     repo,
	}

	if cfg.TLS.Enabled {
		app.autocertMgr = NewAutocertManager(logger, &cfg.TLS)
		tlsConfig := NewTLSConfig(app.autocertMgr)

		app.server = &http.Server{
			Addr:        cfg.TLS.HTTPSPort,
			Handler:     handler,
			TLSConfig:   tlsConfig,
			ReadTimeout: cfg.HTTPServer.Timeouts.Request,
			IdleTimeout: cfg.HTTPServer.Timeouts.Idle,
		}

		app.httpRedirect = &http.Server{
			Addr:    cfg.TLS.HTTPPort,
			Handler: app.autocertMgr.HTTPHandler(createHTTPSRedirectHandler()),
		}

		logger.Info("HTTPS режим активирован", slog.String("domain", cfg.TLS.Domain))
	} else {

		app.server = &http.Server{
			Addr:        cfg.HTTPServer.Address,
			Handler:     handler,
			ReadTimeout: cfg.HTTPServer.Timeouts.Request,
			IdleTimeout: cfg.HTTPServer.Timeouts.Idle,
		}

		logger.Info("HTTP режим (локальная разработка)")
	}

	return app
}

func (app *ServerApp) Run() error {
	const op = "app.Run"
	logger := app.logger.With(slog.String("op", op))

	if app.cfg.TLS.Enabled {
		go func() {
			logger.Info("Запуск HTTP сервера для ACME и редиректов",
				slog.String("address", app.httpRedirect.Addr),
			)
			if err := app.httpRedirect.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Error("HTTP сервер остановлен с ошибкой", slog.String("error", err.Error()))
			}
		}()

		logger.Info("Запуск HTTPS сервера с автоматическими сертификатами",
			slog.String("address", app.server.Addr),
			slog.String("domain", app.cfg.TLS.Domain),
		)

		return app.server.ListenAndServeTLS("", "")
	}

	logger.Info("Запуск HTTP сервера (локальная разработка)",
		slog.String("address", app.server.Addr),
	)
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

	if app.httpRedirect != nil {
		if err := app.httpRedirect.Shutdown(ctx); err != nil {
			logger.Warn("HTTP сервер не успел завершиться", slog.String("error", err.Error()))
			if err := app.httpRedirect.Close(); err != nil {
				logger.Error("Критическая ошибка остановки HTTP сервера", slog.String("error", err.Error()))
			}
		}
	}
}

func registerHandlers(repo *repository.DBRepository, logger *slog.Logger) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

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

	// HTMX endpoints для поиска
	searchHandler := handlers.NewSearchHandler(repo, logger)

	r.Get("/search/filters", searchHandler.SearchFiltersHandler)
	r.Get("/search/results", searchHandler.SearchResultsHandler)

	// TODO Реализовать API endpoints для форм
	// r.Post("/api/contribute/tournament", handlers.ContributeTournament)
	// r.Post("/api/contribute/athlete", handlers.ContributeAthlete)
	// r.Post("/api/contribute/error", handlers.ContributeError)

	return r
}

func createHTTPSRedirectHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		target := "https://" + r.Host + r.URL.Path
		if r.URL.RawQuery != "" {
			target += "?" + r.URL.RawQuery
		}
		http.Redirect(w, r, target, http.StatusMovedPermanently)
	})
}
