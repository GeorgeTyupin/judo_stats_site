package handlers

import (
	"context"
	"strconv"

	"judo_stats_site/internal/handlers/dto"
	"judo_stats_site/templates/components"
	"judo_stats_site/templates/pages"
	"log/slog"
	"net/http"
)

const component = "SearchHandler"

type SearchService interface {
	GeneralSearch(ctx context.Context, query string) ([]any, error)
	JudokaSearch(ctx context.Context, query string, filter dto.JudokaFilters) ([]dto.JudokaResponse, error)
	TournamentSearch(ctx context.Context, query string, filter dto.TournamentFilters) ([]dto.TournamentResponse, error)
	SportClubSearch(ctx context.Context, query string, filter dto.SportClubFilters) ([]dto.SportClubResponse, error)
	CitySearch(ctx context.Context, query string, filter dto.CityFilters) ([]dto.CityResponse, error)
}

type SearchHandler struct {
	service SearchService
	logger  *slog.Logger
}

func NewSearchHandler(service SearchService, logger *slog.Logger) *SearchHandler {
	logger = logger.With(slog.String("component", component))

	return &SearchHandler{
		service: service,
		logger:  logger,
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
		results, err := h.service.GeneralSearch(ctx, query)
		if err != nil {
			h.logger.Error("Ошибка получения результатов поиска", slog.String("error", err.Error()))
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			return
		}

		if len(results) == 0 {
			components.EmptySearchResults().Render(ctx, w)
			return
		}

		// TODO: Реализовать рендеринг смешанных результатов для CategoryAll
		components.EmptySearchResults().Render(ctx, w)

	case dto.CategoryJudoka:
		filters := parseJudokaFilters(r)
		judokas, err := h.service.JudokaSearch(ctx, query, filters)
		if err != nil {
			h.logger.Error("Ошибка получения дзюдоистов", slog.String("error", err.Error()))
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			return
		}

		if len(judokas) == 0 {
			components.EmptySearchResults().Render(ctx, w)
			return
		}

		components.JudokaSearchResults(judokas).Render(ctx, w)

	case dto.CategoryTournament:
		filters := parseTournamentFilters(r)
		tournaments, err := h.service.TournamentSearch(ctx, query, filters)
		if err != nil {
			h.logger.Error("Ошибка получения турниров", slog.String("error", err.Error()))
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			return
		}

		if len(tournaments) == 0 {
			components.EmptySearchResults().Render(ctx, w)
			return
		}

		components.TournamentSearchResults(tournaments).Render(ctx, w)

	case dto.CategorySportClub:
		filters := parseSportClubFilters(r)
		clubs, err := h.service.SportClubSearch(ctx, query, filters)
		if err != nil {
			h.logger.Error("Ошибка получения клубов", slog.String("error", err.Error()))
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			return
		}

		if len(clubs) == 0 {
			components.EmptySearchResults().Render(ctx, w)
			return
		}

		components.SportClubSearchResults(clubs).Render(ctx, w)

	case dto.CategoryCity:
		filters := parseCityFilters(r)
		cities, err := h.service.CitySearch(ctx, query, filters)
		if err != nil {
			h.logger.Error("Ошибка получения городов", slog.String("error", err.Error()))
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			return
		}

		if len(cities) == 0 {
			components.EmptySearchResults().Render(ctx, w)
			return
		}

		components.CitySearchResults(cities).Render(ctx, w)

	default:
		results, err := h.service.GeneralSearch(ctx, query)
		if err != nil {
			h.logger.Error("Ошибка получения результатов поиска", slog.String("error", err.Error()))
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			return
		}

		if len(results) == 0 {
			components.EmptySearchResults().Render(ctx, w)
			return
		}

		// TODO: Реализовать рендеринг смешанных результатов для default
		components.EmptySearchResults().Render(ctx, w)
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
