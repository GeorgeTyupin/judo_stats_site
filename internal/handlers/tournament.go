package handlers

import (
	"judo_stats_site/internal/repository/entity"
	"judo_stats_site/templates/pages"
	"net/http"
)

func Tournament(w http.ResponseWriter, r *http.Request) {
	data := entity.Tournament{
		Name:   "Чемпионат России по дзюдо 1999",
		Type:   "Чемпионат России",
		Date:   "15-18 апреля 1999",
		Place:  "СК «Олимпийский»",
		Year:   1999,
		Month:  4,
		Gender: "Men",
	}

	component := pages.Tournament(data)
	component.Render(r.Context(), w)
}
