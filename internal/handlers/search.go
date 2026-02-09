package handlers

import (
	"judo_stats_site/internal/domain/dto"
	"judo_stats_site/internal/domain/models"
	"judo_stats_site/templates/components"
	"judo_stats_site/templates/pages"
	"log/slog"
	"net/http"
)

const component = "SearchHandler"

type SearchRepository interface {
	// TODO Сделать методы, когда они будут реализованы в репозиторном слое
	GeneralSearch(query string) []any
	JudokaSearch(query string, filter dto.JudokaFilters) []models.Judoka
	TournamentSearch(query string, filter dto.TournamentFilters) []models.Tournament
	SportClubSearch(query string, filter dto.SportClubFilters) []models.SportClub
	CitySearch(query string, filter dto.CityFilters) []models.City
}

type SearchHandler struct {
	repo   SearchRepository
	logger *slog.Logger
}

func NewSearchHandler(repository SearchRepository, logger *slog.Logger) *SearchHandler {
	logger = logger.With(slog.String("component", component))

	return &SearchHandler{
		repo:   repository,
		logger: logger,
	}
}

// SearchFiltersHandler возвращает фильтры для выбранной категории
func (h *SearchHandler) SearchFiltersHandler(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")

	switch dto.SearchCategory(category) {
	case dto.CategoryAll:
		pages.IndexAllFilters().Render(r.Context(), w)
	case dto.CategoryJudoka:
		pages.IndexJudokaFilters().Render(r.Context(), w)
	case dto.CategoryTournament:
		pages.IndexTournamentFilters().Render(r.Context(), w)
	case dto.CategorySportClub:
		pages.IndexSportClubFilters().Render(r.Context(), w)
	case dto.CategoryCity:
		pages.IndexCityFilters().Render(r.Context(), w)
	default:
		pages.IndexAllFilters().Render(r.Context(), w)
	}
}

// SearchResultsHandler возвращает результаты поиска
func (h *SearchHandler) SearchResultsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	category := dto.SearchCategory(r.URL.Query().Get("category"))

	if category == "" {
		category = dto.CategoryAll
	}

	switch dto.SearchCategory(category) {
	case dto.CategoryAll:
		results := h.repo.GeneralSearch(query)

		if len(results) == 0 {
			components.EmptySearchResults().Render(r.Context(), w)
			return
		}

		// TODO: Реализовать рендеринг смешанных результатов для CategoryAll
		components.EmptySearchResults().Render(r.Context(), w)

	case dto.CategoryJudoka:
		filters := parseJudokaFilters(r)
		judokas := h.repo.JudokaSearch(query, filters)

		if len(judokas) == 0 {
			components.EmptySearchResults().Render(r.Context(), w)
			return
		}

		components.JudokaSearchResults(judokas).Render(r.Context(), w)

	case dto.CategoryTournament:
		filters := parseTournamentFilters(r)
		tournaments := h.repo.TournamentSearch(query, filters)

		if len(tournaments) == 0 {
			components.EmptySearchResults().Render(r.Context(), w)
			return
		}

		components.TournamentSearchResults(tournaments).Render(r.Context(), w)

	case dto.CategorySportClub:
		filters := parseSportClubFilters(r)
		clubs := h.repo.SportClubSearch(query, filters)

		if len(clubs) == 0 {
			components.EmptySearchResults().Render(r.Context(), w)
			return
		}

		components.SportClubSearchResults(clubs).Render(r.Context(), w)

	case dto.CategoryCity:
		filters := parseCityFilters(r)
		cities := h.repo.CitySearch(query, filters)

		if len(cities) == 0 {
			components.EmptySearchResults().Render(r.Context(), w)
			return
		}

		components.CitySearchResults(cities).Render(r.Context(), w)

	default:
		results := h.repo.GeneralSearch(query)

		if len(results) == 0 {
			components.EmptySearchResults().Render(r.Context(), w)
			return
		}

		// TODO: Реализовать рендеринг смешанных результатов для default
		components.EmptySearchResults().Render(r.Context(), w)
	}
}

func parseJudokaFilters(r *http.Request) dto.JudokaFilters {
	return dto.JudokaFilters{
		WeightCategory: r.URL.Query().Get("filter_weight"),
		Gender:         r.URL.Query().Get("filter_gender"),
		AgeGroup:       r.URL.Query().Get("filter_age_group"),
		Country:        r.URL.Query().Get("filter_country"),
		SportClub:      r.URL.Query().Get("filter_sportclub"),
		City:           r.URL.Query().Get("filter_city"),
	}
}

func parseTournamentFilters(r *http.Request) dto.TournamentFilters {
	return dto.TournamentFilters{
		Type:     r.URL.Query().Get("filter_type"),
		AgeGroup: r.URL.Query().Get("filter_age_group"),
		Year:     r.URL.Query().Get("filter_year"),
		City:     r.URL.Query().Get("filter_city"),
		Region:   r.URL.Query().Get("filter_region"),
		DateFrom: r.URL.Query().Get("filter_date_from"),
		DateTo:   r.URL.Query().Get("filter_date_to"),
	}
}

func parseSportClubFilters(r *http.Request) dto.SportClubFilters {
	return dto.SportClubFilters{
		City:     r.URL.Query().Get("filter_city"),
		Region:   r.URL.Query().Get("filter_region"),
		Founded:  r.URL.Query().Get("filter_founded"),
		HasCoach: r.URL.Query().Get("filter_has_coach") == "true",
	}
}
func parseCityFilters(r *http.Request) dto.CityFilters {
	return dto.CityFilters{
		Region: r.URL.Query().Get("filter_region"),
	}
}
