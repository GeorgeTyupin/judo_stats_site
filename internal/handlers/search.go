package handlers

import (
	"judo_stats_site/internal/domain/dto"
	"judo_stats_site/templates/components"
	"judo_stats_site/templates/pages"
	"net/http"
)

// SearchFiltersHandler возвращает фильтры для выбранной категории
func SearchFiltersHandler(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")

	switch dto.SearchCategory(category) {
	case dto.CategoryJudoka:
		pages.IndexJudokaFilters().Render(r.Context(), w)
	case dto.CategoryTournament:
		pages.IndexTournamentFilters().Render(r.Context(), w)
	case dto.CategorySportClub:
		pages.IndexSportClubFilters().Render(r.Context(), w)
	case dto.CategoryCity:
		pages.IndexCityFilters().Render(r.Context(), w)
	default:
		pages.IndexJudokaFilters().Render(r.Context(), w)
	}
}

// SearchResultsHandler возвращает результаты поиска
func SearchResultsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO сделать обработку поиска и форм с фильтрами
	// query := r.URL.Query().Get("query")
	// category := dto.SearchCategory(r.URL.Query().Get("category"))

	// if category == "" {

	// }

	components.EmptySearchResults().Render(r.Context(), w)
}
