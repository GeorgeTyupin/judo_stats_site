package handlers

import (
	"judo_stats_site/internal/repository/record"
	"judo_stats_site/templates/pages"
	"net/http"
)

func City(w http.ResponseWriter, r *http.Request) {
	// Пример данных - позже это будет из БД
	data := record.City{
		Name:        "Москва",
		Description: "Москва - столица России и один из крупнейших центров развития дзюдо в стране. В городе действуют десятки спортивных клубов и секций, которые воспитывают чемпионов всероссийского и международного уровня.",
	}

	component := pages.City(data)
	component.Render(r.Context(), w)
}
