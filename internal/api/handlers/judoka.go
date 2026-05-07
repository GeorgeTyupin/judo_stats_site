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

type JudokaService interface {
	GetJudokaByID(ctx context.Context, id int64) (record.Judoka, error)
}

type JudokaHandler struct {
	logger  *slog.Logger
	service JudokaService
}

func NewJudokaHandler(service JudokaService, logger *slog.Logger) *JudokaHandler {
	return &JudokaHandler{logger: logger, service: service}
}

func (h *JudokaHandler) Page(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.JudokaHandler.Page"
	logger := h.logger.With(slog.String("op", op))

	data := record.Judoka{
		LastName:         "Emelianenko",
		FirstName:        "Fedor",
		LastNameRus:      "Емельяненко",
		FirstNameRus:     "Федор",
		Country:          "Россия",
		WeightCategories: []string{"+100 кг"},
		BirthDate:        "28 сентября 1976",
		BirthPlace:       "Рубежное, Луганская обл.",
	}
	httputil.Render(r.Context(), w, logger, pages.Judoka(data))
}

func (h *JudokaHandler) PageByID(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.JudokaHandler.PageByID"
	logger := h.logger.With(slog.String("op", op))

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	data, err := h.service.GetJudokaByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.NotFound(w, r)
			return
		}
		logger.Error("ошибка получения дзюдоиста", slog.Int64("id", id), slog.String("error", err.Error()))
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	httputil.Render(r.Context(), w, logger, pages.Judoka(data))
}
