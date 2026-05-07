package handlers

import (
	"judo_stats_site/internal/api/httputil"
	"judo_stats_site/templates/pages"
	"log/slog"
	"net/http"
)

type IndexHandler struct {
	logger *slog.Logger
}

func NewIndexHandler(logger *slog.Logger) *IndexHandler {
	return &IndexHandler{logger: logger}
}

func (h *IndexHandler) Page(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.IndexHandler.Page"
	logger := h.logger.With(slog.String("op", op))

	httputil.Render(r.Context(), w, logger, pages.Index())
}
