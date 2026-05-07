package handlers

import (
	"judo_stats_site/internal/api/httputil"
	"judo_stats_site/templates/pages"
	"log/slog"
	"net/http"
)

type ContributeHandler struct {
	logger *slog.Logger
}

func NewContributeHandler(logger *slog.Logger) *ContributeHandler {
	return &ContributeHandler{logger: logger}
}

func (h *ContributeHandler) Page(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.ContributeHandler.Page"
	logger := h.logger.With(slog.String("op", op))

	httputil.Render(r.Context(), w, logger, pages.Contribute())
}
