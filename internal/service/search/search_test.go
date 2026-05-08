package search

import (
	"context"
	"judo_stats_site/internal/api/handlers/dto"
	"judo_stats_site/internal/repository/record"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockSearchRepository struct{}

func (m *mockSearchRepository) GeneralSearch(ctx context.Context, query string) ([]any, error) {
	return nil, nil
}
func (m *mockSearchRepository) JudokaSearch(ctx context.Context, query string, filter dto.JudokaFilters) ([]record.Judoka, error) {
	if query == "" {
		return []record.Judoka{}, nil
	}
	return []record.Judoka{
		{LastNameRus: query},
	}, nil
}
func (m *mockSearchRepository) GetJudokaByID(ctx context.Context, id int64) (record.Judoka, error) {
	return record.Judoka{}, nil
}
func (m *mockSearchRepository) GetTournamentByID(ctx context.Context, id int64) (record.Tournament, error) {
	return record.Tournament{}, nil
}
func (m *mockSearchRepository) TournamentSearch(ctx context.Context, query string, filter dto.TournamentFilters) ([]record.Tournament, error) {
	return nil, nil
}
func (m *mockSearchRepository) SportClubSearch(ctx context.Context, query string, filter dto.SportClubFilters) ([]record.SportClub, error) {
	return nil, nil
}
func (m *mockSearchRepository) CitySearch(ctx context.Context, query string, filter dto.CityFilters) ([]record.City, error) {
	return nil, nil
}

func newTestSearchService() *SearchService {
	return &SearchService{
		repo: &mockSearchRepository{},
	}
}

func TestJudokaSearch(t *testing.T) {
	service := newTestSearchService()

	data := []struct {
		name  string
		query string
		want  []dto.JudokaResponse
	}{
		{
			name:  "Пустой запрос",
			query: "",
			want:  []dto.JudokaResponse{},
		},
		{
			name:  "Канонизация",
			query: "Семён",
			want:  []dto.JudokaResponse{{Name: "Семен"}},
		},
		{
			name:  "Обычный поиск",
			query: "Иван",
			want:  []dto.JudokaResponse{{Name: "Иван"}},
		},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.JudokaSearch(context.Background(), tt.query, dto.JudokaFilters{})
			require.NoError(t, err)

			assert.Equal(t, tt.want, got)
		})
	}

}
