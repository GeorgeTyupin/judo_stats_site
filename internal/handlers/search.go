package handlers

import (
	"judo_stats_site/internal/models"
	"judo_stats_site/templates/components"
	"judo_stats_site/templates/pages"
	"net/http"
)

// SearchFiltersHandler возвращает фильтры для выбранной категории
func SearchFiltersHandler(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")

	switch models.SearchCategory(category) {
	case models.CategoryJudoka:
		pages.IndexJudokaFilters().Render(r.Context(), w)
	case models.CategoryTournament:
		pages.IndexTournamentFilters().Render(r.Context(), w)
	case models.CategorySportClub:
		pages.IndexSportClubFilters().Render(r.Context(), w)
	case models.CategoryCity:
		pages.IndexCityFilters().Render(r.Context(), w)
	default:
		pages.IndexJudokaFilters().Render(r.Context(), w)
	}
}

// SearchResultsHandler возвращает результаты поиска
func SearchResultsHandler(w http.ResponseWriter, r *http.Request) {
	// Пока просто возвращаем пустые результаты
	// TODO: Реализовать логику поиска на бэкенде
	components.EmptySearchResults().Render(r.Context(), w)
}
