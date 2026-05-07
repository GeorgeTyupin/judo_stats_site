package handlers

import (
	"judo_stats_site/internal/api/httputil"
	"judo_stats_site/internal/repository/record"
	"judo_stats_site/templates/pages"
	"log/slog"
	"net/http"
)

type CityHandler struct {
	logger *slog.Logger
}

func NewCityHandler(logger *slog.Logger) *CityHandler {
	return &CityHandler{logger: logger}
}

func (h *CityHandler) Page(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.CityHandler.Page"
	logger := h.logger.With(slog.String("op", op))

	// Пример данных - позже это будет из БД
	data := record.City{
		Name:        "Москва",
		Description: "Москва - столица России и один из крупнейших центров развития дзюдо в стране. В городе действуют десятки спортивных клубов и секций, которые воспитывают чемпионов всероссийского и международного уровня.",
	}
	httputil.Render(r.Context(), w, logger, pages.City(data))
}
