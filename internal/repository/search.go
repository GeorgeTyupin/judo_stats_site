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

// judokaSortFields — whitelist допустимых полей сортировки для дзюдоистов
var judokaSortFields = map[string]string{
	"name":       "j.last_name",
	"country":    "ca.country",
	"birth_date": "j.birth_date",
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
	sortCol := resolveSortField(judokaSortFields, filter.SortBy, "j.last_name")
	sortDir := "ASC"
	if filter.SortDir == "desc" {
		sortDir = "DESC"
	}

	sqlQuery := fmt.Sprintf(`
		SELECT
			j.id, j.last_name, j.first_name,
			COALESCE(j.last_name_rus,  '') AS last_name_rus,
			COALESCE(j.first_name_rus, '') AS first_name_rus,
			COALESCE(j.gender,         '') AS gender,
			COALESCE(j.weight_category, '{}'::text[]) AS weight_category,
			COALESCE(j.birth_date,     '') AS birth_date,
			COALESCE(j.birth_place,    '') AS birth_place,
			COALESCE(ca.country,       '') AS country,
			COALESCE(ci.city,          '') AS city,
			COALESCE(clubs.sport_club,  '') AS sport_club
		FROM judokas j
		LEFT JOIN (
			SELECT jco.judoka_id, STRING_AGG(c.name, ', ') AS country
			FROM judoka_countries jco
			JOIN countries c ON jco.country_id = c.id
			GROUP BY jco.judoka_id
		) ca ON j.id = ca.judoka_id
		LEFT JOIN (
			SELECT jcit.judoka_id, STRING_AGG(ct.name, ', ') AS city
			FROM judoka_cities jcit
			JOIN cities ct ON jcit.city_id = ct.id
			GROUP BY jcit.judoka_id
		) ci ON j.id = ci.judoka_id
		LEFT JOIN (
			SELECT jsc.judoka_id, STRING_AGG(sclub.name, ', ') AS sport_club
			FROM judoka_sport_clubs jsc
			JOIN sport_clubs sclub ON jsc.sport_club_id = sclub.id
			GROUP BY jsc.judoka_id
		) clubs ON j.id = clubs.judoka_id
		WHERE
			concat_ws(' ', j.last_name, j.first_name, j.last_name_rus, j.first_name_rus)
				ILIKE '%%' || $1 || '%%'
		AND ($2 = '' OR j.gender = $2)
		AND ($3 = '' OR j.birth_date LIKE '%%' || $3 || '%%')
		AND ($4 = '' OR ca.country      ILIKE '%%' || $4 || '%%')
		AND ($5 = '' OR clubs.sport_club ILIKE '%%' || $5 || '%%')
		AND ($6 = '' OR ci.city         ILIKE '%%' || $6 || '%%')
		AND (coalesce(cardinality($7::text[]), 0) = 0 OR j.weight_category && $7::text[])
		ORDER BY %s %s
		LIMIT %d`,
		sortCol, sortDir, r.selectRowsLimit,
	)

	rows, err := r.db.Query(ctx, sqlQuery,
		query,
		filter.Gender,
		filter.BirthYear,
		filter.Country,
		filter.SportClub,
		filter.City,
		filter.WeightCategories,
	)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса в бд: %w", err)
	}

	judokas, err := pgx.CollectRows(rows, pgx.RowToStructByName[record.Judoka])
	if err != nil {
		return nil, fmt.Errorf("ошибка преобразования rows к record.Judoka: %w", err)
	}

	return judokas, nil
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
