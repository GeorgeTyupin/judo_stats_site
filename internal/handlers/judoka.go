package handlers

import (
	"judo_stats_site/templates/pages"
	"net/http"
)

func Judoka(w http.ResponseWriter, r *http.Request) {
	// Пример данных - позже это будет из БД
	data := pages.JudokaData{
		Name:           "Федор Емельяненко",
		Country:        "Россия",
		CountryFlag:    "🇷🇺",
		WeightCategory: "+100 кг",
		BirthDate:      "28 сентября 1976",
		Age:            47,
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
