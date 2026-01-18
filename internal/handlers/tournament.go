package handlers

import (
	"judo_stats_site/internal/models"
	"judo_stats_site/templates/pages"
	"net/http"
)

func Tournament(w http.ResponseWriter, r *http.Request) {
	// Пример данных - позже это будет из БД
	data := models.TournamentData{
		Name:         "Чемпионат России по дзюдо 1999",
		Type:         "Чемпионат России",
		Date:         "15-18 апреля 1999",
		Location:     "Москва, Россия",
		City:         "Москва",
		Participants: 245,
		Categories:   14,
		Organizer:    "Федерация дзюдо России",
		Venue:        "СК «Олимпийский»",
		AgeGroup:     "Взрослые (18+)",
		System:       "Олимпийская система с утешительными поединками",
		Regions:      54,
		Matches:      487,
	}

	component := pages.Tournament(data)
	component.Render(r.Context(), w)
}
