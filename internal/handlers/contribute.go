package handlers

import (
	"judo_stats_site/templates/pages"
	"net/http"
)

func Contribute(w http.ResponseWriter, r *http.Request) {
	component := pages.Contribute()
	component.Render(r.Context(), w)
}
