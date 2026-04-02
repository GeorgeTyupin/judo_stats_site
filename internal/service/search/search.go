package search

import (
	"context"
	"fmt"
	"log/slog"

	"judo_stats_site/internal/api/handlers/dto"
	"judo_stats_site/internal/repository/record"
	"judo_stats_site/pkg/lib/canonize"
)

// SearchRepository определяет методы работы с базой данных для поиска
type SearchRepository interface {
	GeneralSearch(ctx context.Context, query string) ([]any, error)
	JudokaSearch(ctx context.Context, query string, filter dto.JudokaFilters) ([]record.Judoka, error)
	TournamentSearch(ctx context.Context, query string, filter dto.TournamentFilters) ([]record.Tournament, error)
	SportClubSearch(ctx context.Context, query string, filter dto.SportClubFilters) ([]record.SportClub, error)
	CitySearch(ctx context.Context, query string, filter dto.CityFilters) ([]record.City, error)
	GetJudokaByID(ctx context.Context, id int64) (record.Judoka, error)
	GetTournamentByID(ctx context.Context, id int64) (record.Tournament, error)
}

type SearchService struct {
	repo   SearchRepository
	logger *slog.Logger
}

func NewSearchService(repo SearchRepository, logger *slog.Logger) *SearchService {
	return &SearchService{
		repo:   repo,
		logger: logger,
	}
}

func (s *SearchService) GeneralSearch(ctx context.Context, query string) ([]any, error) {
	return s.repo.GeneralSearch(ctx, query)
}

func (s *SearchService) JudokaSearch(ctx context.Context, query string, filter dto.JudokaFilters) ([]dto.JudokaResponse, error) {
	query, err := canonize.Canonize(query)
	if err != nil {
		return nil, fmt.Errorf("не удалось канонизировать запрос: %w", err)
	}

	judokas, err := s.repo.JudokaSearch(ctx, query, filter)
	if err != nil {
		return nil, fmt.Errorf("не удалось выполнить поиск: %w", err)
	}

	responses := make([]dto.JudokaResponse, len(judokas))
	for i, j := range judokas {
		responses[i] = dto.JudokaResponse{
			ID:               j.ID,
			Name:             j.LastNameRus + " " + j.FirstNameRus,
			Country:          j.Country,
			City:             j.City,
			SportClub:        j.SportClub,
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
			ID:       c.ID,
			Name:     c.Name,
			FullName: c.FullName,
		}
	}
	return responses, nil
}

func (s *SearchService) GetJudokaByID(ctx context.Context, id int64) (record.Judoka, error) {
	return s.repo.GetJudokaByID(ctx, id)
}

func (s *SearchService) GetTournamentByID(ctx context.Context, id int64) (record.Tournament, error) {
	return s.repo.GetTournamentByID(ctx, id)
}

func (s *SearchService) CitySearch(ctx context.Context, query string, filter dto.CityFilters) ([]dto.CityResponse, error) {
	cities, err := s.repo.CitySearch(ctx, query, filter)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.CityResponse, len(cities))
	for i, c := range cities {
		responses[i] = dto.CityResponse{
			ID:          c.ID,
			Name:        c.Name,
			Description: c.Description,
		}
	}
	return responses, nil
}
