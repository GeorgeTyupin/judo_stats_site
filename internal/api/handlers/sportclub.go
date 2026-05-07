package handlers

import (
	"judo_stats_site/internal/api/httputil"
	"judo_stats_site/internal/repository/record"
	"judo_stats_site/templates/pages"
	"log/slog"
	"net/http"
)

type SportClubHandler struct {
	logger *slog.Logger
}

func NewSportClubHandler(logger *slog.Logger) *SportClubHandler {
	return &SportClubHandler{logger: logger}
}

func (h *SportClubHandler) Page(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.SportClubHandler.Page"
	logger := h.logger.With(slog.String("op", op))

	data := record.SportClub{
		Name:        "ЦСКА",
		FullName:    "Центральный спортивный клуб армии",
		Founded:     "29 апреля 1923",
		Region:      "Московская область",
		HeadCoach:   "Иванов И.И.",
		Description: "ЦСКА - одно из старейших и наиболее титулованных спортивных обществ России. За свою историю клуб воспитал множество чемпионов и призеров крупнейших международных турниров.",
	}
	httputil.Render(r.Context(), w, logger, pages.SportClub(data))
}
