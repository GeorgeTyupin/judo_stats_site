package handlers

import (
	"judo_stats_site/internal/models"
	"judo_stats_site/templates/pages"
	"net/http"
)

func Judoka(w http.ResponseWriter, r *http.Request) {
	data := models.Judoka{
		Name:           "Федор Емельяненко",
		Country:        "Россия",
		CountryFlag:    "🇷🇺",
		WeightCategory: "+100 кг",
		BirthDate:      "28 сентября 1976",
		BirthPlace:     "Рубежное, Луганская обл.",
		SportClub:      "ЦСКА",
		Coach:          "Бородин В.А.",
		GoldMedals:     5,
		SilverMedals:   3,
		BronzeMedals:   2,
	}

	component := pages.Judoka(data)
	component.Render(r.Context(), w)
}
