package handlers

import (
	"judo_stats_site/internal/repository/record"
	"judo_stats_site/templates/pages"
	"net/http"
)

func SportClub(w http.ResponseWriter, r *http.Request) {
	data := record.SportClub{
		Name:        "ЦСКА",
		FullName:    "Центральный спортивный клуб армии",
		Founded:     "29 апреля 1923",
		Region:      "Московская область",
		HeadCoach:   "Иванов И.И.",
		Description: "ЦСКА - одно из старейших и наиболее титулованных спортивных обществ России. За свою историю клуб воспитал множество чемпионов и призеров крупнейших международных турниров.",
	}

	component := pages.SportClub(data)
	component.Render(r.Context(), w)
}
