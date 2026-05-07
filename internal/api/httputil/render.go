package httputil

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

func Render(ctx context.Context, w http.ResponseWriter, logger *slog.Logger, c templ.Component) {
	if err := c.Render(ctx, w); err != nil {
		logger.Error("ошибка рендеринга", slog.String("error", err.Error()))
	}
}
