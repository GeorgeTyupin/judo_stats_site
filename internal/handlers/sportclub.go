package handlers

import (
	"judo_stats_site/internal/models"
	"judo_stats_site/templates/pages"
	"net/http"
)

func SportClub(w http.ResponseWriter, r *http.Request) {
	data := models.SportClub{
		Name:           "ЦСКА",
		FullName:       "Центральный спортивный клуб армии",
		Founded:        "29 апреля 1923",
		Region:         "Московская область",
		HeadCoach:      "Иванов И.И.",
		Athletes:       156,
		ActiveAthletes: 98,
		Coaches:        12,
		GoldMedals:     245,
		SilverMedals:   189,
		BronzeMedals:   156,
		TotalMedals:    590,
		Description:    "ЦСКА - одно из старейших и наиболее титулованных спортивных обществ России. За свою историю клуб воспитал множество чемпионов и призеров крупнейших международных турниров.",
	}

	component := pages.SportClub(data)
	component.Render(r.Context(), w)
}
