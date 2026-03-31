package repository

import (
	"context"
	"fmt"
	"judo_stats_site/internal/api/handlers/dto"
	"judo_stats_site/internal/config"
	"judo_stats_site/internal/repository/record"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SearchRepository struct {
	db              *pgxpool.Pool
	logger          *slog.Logger
	timeout         time.Duration
	selectRowsLimit int32
}

func NewSearchRepository(db *pgxpool.Pool, logger *slog.Logger, cfg *config.Database) *SearchRepository {
	return &SearchRepository{
		db:              db,
		logger:          logger,
		timeout:         cfg.Timeout,
		selectRowsLimit: cfg.SelectRowsLimit,
	}
}

// tournamentSortFields — whitelist допустимых полей сортировки для турниров
var tournamentSortFields = map[string]string{
	"name":  "tournaments.name",
	"type":  "tournaments.type",
	"place": "tournaments.place",
	"date":  "tournaments.year",
	"year":  "tournaments.year",
}

func resolveSortField(fields map[string]string, sortBy, defaultField string) string {
	if col, ok := fields[sortBy]; ok {
		return col
	}
	return defaultField
}

func (r *SearchRepository) GeneralSearch(ctx context.Context, query string) ([]any, error) {
	// TODO Реализовать метод для общего поиска
	return nil, nil
}

func (r *SearchRepository) JudokaSearch(ctx context.Context, query string, filter dto.JudokaFilters) ([]record.Judoka, error) {
	// TODO: Реализовать поиск дзюдоистов в БД
	return nil, nil
}

func (r *SearchRepository) TournamentSearch(ctx context.Context, query string, filter dto.TournamentFilters) ([]record.Tournament, error) {
	sortCol := resolveSortField(tournamentSortFields, filter.SortBy, "tournaments.name")
	sortDir := "ASC"
	if filter.SortDir == "desc" {
		sortDir = "DESC"
	}

	sqlQuery := fmt.Sprintf(`
		SELECT tournaments.* FROM tournaments
		LEFT JOIN cities ON tournaments.city_id = cities.id
		LEFT JOIN countries ON tournaments.country_id = countries.id
		LEFT JOIN ussr_republics ON tournaments.republic_id = ussr_republics.id
		WHERE tournaments.name LIKE '%%' || $1 || '%%'
		AND ($2 = '' OR tournaments.type = $2)
		AND ($3 = '' OR tournaments.gender = $3)
		AND ($4 = 0 OR tournaments.year = $4)
		AND ($5 = 0 OR tournaments.month = $5)
		AND ($6 = '' OR cities.name ILIKE '%%' || $6 || '%%')
		AND ($7 = '' OR ussr_republics.name ILIKE '%%' || $7 || '%%')
		AND ($8 = '' OR countries.name ILIKE '%%' || $8 || '%%')
		ORDER BY %s %s
		LIMIT %d`,
		sortCol, sortDir, r.selectRowsLimit,
	)

	rows, err := r.db.Query(ctx, sqlQuery,
		query,
		filter.Type,
		filter.Gender,
		filter.Year,
		filter.Month,
		filter.City,
		filter.Republic,
		filter.Country,
	)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса в бд: %w", err)
	}

	tournaments, err := pgx.CollectRows(rows, pgx.RowToStructByName[record.Tournament])
	if err != nil {
		return nil, fmt.Errorf("ошибка преобразования rows к record.Tournament: %w", err)
	}

	return tournaments, nil
}

func (r *SearchRepository) SportClubSearch(ctx context.Context, query string, filter dto.SportClubFilters) ([]record.SportClub, error) {
	// TODO Реализовать метод для поиска СО
	return nil, nil
}

func (r *SearchRepository) CitySearch(ctx context.Context, query string, filter dto.CityFilters) ([]record.City, error) {
	// TODO Реализовать метод для поиска города
	return nil, nil
}
