package handlers

import (
	"judo_stats_site/internal/api/httputil"
	"judo_stats_site/internal/repository/record"
	"judo_stats_site/templates/pages"
	"log/slog"
	"net/http"
)

type JudokaHandler struct {
	logger *slog.Logger
}

func NewJudokaHandler(logger *slog.Logger) *JudokaHandler {
	return &JudokaHandler{logger: logger}
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
