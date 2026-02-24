package service

import (
	"context"
	"log/slog"

	"judo_stats_site/internal/handlers/dto"
	"judo_stats_site/internal/repository/entity"
)

// SearchRepository определяет методы работы с базой данных для поиска
type SearchRepository interface {
	GeneralSearch(ctx context.Context, query string) ([]any, error)
	JudokaSearch(ctx context.Context, query string, filter dto.JudokaFilters) ([]entity.Judoka, error)
	TournamentSearch(ctx context.Context, query string, filter dto.TournamentFilters) ([]entity.Tournament, error)
	SportClubSearch(ctx context.Context, query string, filter dto.SportClubFilters) ([]entity.SportClub, error)
	CitySearch(ctx context.Context, query string, filter dto.CityFilters) ([]entity.City, error)
}

type SearchService struct {
	repo   SearchRepository
	logger *slog.Logger
}

func NewSearchService(repo SearchRepository, logger *slog.Logger) *SearchService {
	return &SearchService{
		repo:   repo,
		logger: logger.With(slog.String("component", "SearchService")),
	}
}

func (s *SearchService) GeneralSearch(ctx context.Context, query string) ([]any, error) {
	return s.repo.GeneralSearch(ctx, query)
}

func (s *SearchService) JudokaSearch(ctx context.Context, query string, filter dto.JudokaFilters) ([]dto.JudokaResponse, error) {
	judokas, err := s.repo.JudokaSearch(ctx, query, filter)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.JudokaResponse, len(judokas))
	for i, j := range judokas {
		responses[i] = dto.JudokaResponse{
			ID:               j.ID,
			Name:             j.Name,
			Country:          j.Country,
			BirthPlace:       j.BirthPlace,
			WeightCategories: j.WeightCategories,
			BirthDate:        j.BirthDate,
		}
	}
	return responses, nil
}

func (s *SearchService) TournamentSearch(ctx context.Context, query string, filter dto.TournamentFilters) ([]dto.TournamentResponse, error) {
	tournaments, err := s.repo.TournamentSearch(ctx, query, filter)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.TournamentResponse, len(tournaments))
	for i, t := range tournaments {
		responses[i] = dto.TournamentResponse{
			ID:     t.ID,
			Name:   t.Name,
			Type:   t.Type,
			Place:  t.Place,
			Date:   t.Date,
			Gender: t.Gender,
		}
	}
	return responses, nil
}

func (s *SearchService) SportClubSearch(ctx context.Context, query string, filter dto.SportClubFilters) ([]dto.SportClubResponse, error) {
	clubs, err := s.repo.SportClubSearch(ctx, query, filter)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.SportClubResponse, len(clubs))
	for i, c := range clubs {
		responses[i] = dto.SportClubResponse{
			ID:       int64(c.ID),
			Name:     c.Name,
			FullName: c.FullName,
			Region:   c.Region,
			Founded:  c.Founded,
		}
	}
	return responses, nil
}

func (s *SearchService) CitySearch(ctx context.Context, query string, filter dto.CityFilters) ([]dto.CityResponse, error) {
	cities, err := s.repo.CitySearch(ctx, query, filter)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.CityResponse, len(cities))
	for i, c := range cities {
		responses[i] = dto.CityResponse{
			ID:          int64(c.ID),
			Name:        c.Name,
			Description: c.Description,
		}
	}
	return responses, nil
}
