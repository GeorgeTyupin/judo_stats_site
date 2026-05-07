package handlers

import (
	"context"
	"errors"
	"judo_stats_site/internal/api/httputil"
	"judo_stats_site/internal/repository/record"
	"judo_stats_site/templates/pages"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

type TournamentService interface {
	GetTournamentByID(ctx context.Context, id int64) (record.Tournament, error)
}

type TournamentHandler struct {
	logger  *slog.Logger
	service TournamentService
}

func NewTournamentHandler(service TournamentService, logger *slog.Logger) *TournamentHandler {
	return &TournamentHandler{logger: logger, service: service}
}

func (h *TournamentHandler) Page(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.TournamentHandler.Page"
	logger := h.logger.With(slog.String("op", op))

	data := record.Tournament{
		Name:   "Чемпионат России по дзюдо 1999",
		Type:   "Чемпионат России",
		Date:   "15-18 апреля 1999",
		Place:  "СК «Олимпийский»",
		Year:   1999,
		Month:  4,
		Gender: "Men",
	}
	httputil.Render(r.Context(), w, logger, pages.Tournament(data))
}

func (h *TournamentHandler) PageByID(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.TournamentHandler.PageByID"
	logger := h.logger.With(slog.String("op", op))

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	data, err := h.service.GetTournamentByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.NotFound(w, r)
			return
		}
		logger.Error("ошибка получения турнира", slog.Int64("id", id), slog.String("error", err.Error()))
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	httputil.Render(r.Context(), w, logger, pages.Tournament(data))
}
