package handlers

import (
	"judo_stats_site/internal/repository/entity"
	"judo_stats_site/templates/pages"
	"net/http"
)

func City(w http.ResponseWriter, r *http.Request) {
	// Пример данных - позже это будет из БД
	data := entity.City{
		Name:           "Москва",
		Population:     12655050,
		Founded:        "1147 год",
		SportClubs:     45,
		Athletes:       3500,
		ActiveAthletes: 2100,
		Coaches:        280,
		GoldMedals:     1250,
		SilverMedals:   980,
		BronzeMedals:   756,
		TotalMedals:    2986,
		Description:    "Москва - столица России и один из крупнейших центров развития дзюдо в стране. В городе действуют десятки спортивных клубов и секций, которые воспитывают чемпионов всероссийского и международного уровня.",
	}

	component := pages.City(data)
	component.Render(r.Context(), w)
}
