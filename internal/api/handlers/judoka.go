package handlers

import (
	"judo_stats_site/internal/repository/record"
	"judo_stats_site/templates/pages"
	"net/http"
)

func Judoka(w http.ResponseWriter, r *http.Request) {
	data := record.Judoka{
		LastName:         "Emelianenko",
		FirstName:        "Fedor",
		LastNameRus:      "Емельяненко",
		FirstNameRus:     "Федор",
		Country:          "Россия",
		WeightCategories: []string{"+100 кг"},
		BirthDate:        "28 сентября 1976",
		BirthPlace:       "Рубежное, Луганская обл.",
	}

	component := pages.Judoka(data)
	component.Render(r.Context(), w)
}
