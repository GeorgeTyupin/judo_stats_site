package handlers

import (
	"judo_stats_site/templates/pages"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	component := pages.Index()
	component.Render(r.Context(), w)
}
