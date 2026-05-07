package handlers

import (
	"judo_stats_site/internal/api/httputil"
	"judo_stats_site/internal/repository/record"
	"judo_stats_site/templates/pages"
	"log/slog"
	"net/http"
)

type TournamentHandler struct {
	logger *slog.Logger
}

func NewTournamentHandler(logger *slog.Logger) *TournamentHandler {
	return &TournamentHandler{logger: logger}
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
