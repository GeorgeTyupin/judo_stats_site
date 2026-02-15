package handlers

import (
	"context"
	"strconv"

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
	GeneralSearch(ctx context.Context, query string) ([]any, error)
	JudokaSearch(ctx context.Context, query string, filter dto.JudokaFilters) ([]models.Judoka, error)
	TournamentSearch(ctx context.Context, query string, filter dto.TournamentFilters) ([]models.Tournament, error)
	SportClubSearch(ctx context.Context, query string, filter dto.SportClubFilters) ([]models.SportClub, error)
	CitySearch(ctx context.Context, query string, filter dto.CityFilters) ([]models.City, error)
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
	ctx := r.Context()

	if category == "" {
		category = dto.CategoryAll
	}

	switch dto.SearchCategory(category) {
	case dto.CategoryAll:
		results, _ := h.repo.GeneralSearch(ctx, query)

		if len(results) == 0 {
			components.EmptySearchResults().Render(r.Context(), w)
			return
		}

		// TODO: Реализовать рендеринг смешанных результатов для CategoryAll
		components.EmptySearchResults().Render(r.Context(), w)

	case dto.CategoryJudoka:
		filters := parseJudokaFilters(r)
		judokas, _ := h.repo.JudokaSearch(ctx, query, filters)

		if len(judokas) == 0 {
			components.EmptySearchResults().Render(r.Context(), w)
			return
		}

		components.JudokaSearchResults(judokas).Render(r.Context(), w)

	case dto.CategoryTournament:
		filters := parseTournamentFilters(r)
		tournaments, err := h.repo.TournamentSearch(ctx, query, filters)
		if err != nil {
			h.logger.Error("Ошибка получения турниров", slog.String("error", err.Error()))
		}

		if len(tournaments) == 0 {
			components.EmptySearchResults().Render(r.Context(), w)
			return
		}

		components.TournamentSearchResults(tournaments).Render(r.Context(), w)

	case dto.CategorySportClub:
		filters := parseSportClubFilters(r)
		clubs, _ := h.repo.SportClubSearch(ctx, query, filters)

		if len(clubs) == 0 {
			components.EmptySearchResults().Render(r.Context(), w)
			return
		}

		components.SportClubSearchResults(clubs).Render(r.Context(), w)

	case dto.CategoryCity:
		filters := parseCityFilters(r)
		cities, _ := h.repo.CitySearch(ctx, query, filters)

		if len(cities) == 0 {
			components.EmptySearchResults().Render(r.Context(), w)
			return
		}

		components.CitySearchResults(cities).Render(r.Context(), w)

	default:
		results, _ := h.repo.GeneralSearch(ctx, query)

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
	filters := dto.TournamentFilters{
		Type:     r.URL.Query().Get("filter_type"),
		Gender:   r.URL.Query().Get("filter_gender"),
		City:     r.URL.Query().Get("filter_city"),
		Country:  r.URL.Query().Get("filter_country"),
		Republic: r.URL.Query().Get("filter_republic"),
	}

	year, err := strconv.Atoi(r.URL.Query().Get("filter_year"))
	if err != nil {
		filters.Year = 0
	} else {
		filters.Year = int16(year)
	}

	return filters
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
		Republic: r.URL.Query().Get("filter_republic"),
	}
}
