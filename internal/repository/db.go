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

const component = "Repository"

type DBRepository struct {
	timeout time.Duration
	db      *pgxpool.Pool
	logger  *slog.Logger
}

func NewDBRepository(ctx context.Context, dbConf *config.Database, logger *slog.Logger) (*DBRepository, error) {
	logger = logger.With(slog.String("component", component))

	poolConf, err := pgxpool.ParseConfig(dbConf.GetConnString())
	if err != nil {
		return nil, fmt.Errorf("не удалось создать конфиг для pgxpool. Возникла ошибка %w", err)
	}

	poolConf.MaxConns = dbConf.MaxConns
	poolConf.MinConns = dbConf.MinConns
	poolConf.MaxConnIdleTime = dbConf.MaxConnIdleTime
	poolConf.MaxConnLifetime = dbConf.MaxConnLifetime

	dbPool, err := pgxpool.NewWithConfig(ctx, poolConf)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать пулл базы данных. Возникла ошибка %w", err)
	}

	if err := dbPool.Ping(ctx); err != nil {
		dbPool.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return &DBRepository{
		timeout: dbConf.Timeout,
		db:      dbPool,
		logger:  logger,
	}, nil
}

func (r *DBRepository) Close() {
	const op = "DBRepository.Close"
	logger := r.logger.With(slog.String("op", op))

	r.db.Close()
	logger.Info("Пул подключений к БД закрыт")
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

func (r *DBRepository) GeneralSearch(ctx context.Context, query string) ([]any, error) {
	// TODO Реализовать метод для общего поиска
	return nil, nil
}

func (r *DBRepository) JudokaSearch(ctx context.Context, query string, filter dto.JudokaFilters) ([]record.Judoka, error) {
	// TODO: Реализовать поиск дзюдоистов в БД
	return nil, nil
}

func (r *DBRepository) TournamentSearch(ctx context.Context, query string, filter dto.TournamentFilters) ([]record.Tournament, error) {
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
		AND ($4 = '' OR tournaments.type = $4)
		AND ($5 = 0 OR tournaments.year = $5)
		AND ($6 = 0 OR tournaments.month = $6)
		AND ($7 = '' OR cities.name ILIKE '%%' || $7 || '%%')
		AND ($8 = '' OR ussr_republics.name ILIKE '%%' || $8 || '%%')
		AND ($9 = '' OR countries.name ILIKE '%%' || $9 || '%%')
		ORDER BY %s %s`, sortCol, sortDir)

	rows, err := r.db.Query(ctx, sqlQuery,
		query,
		filter.Type,
		filter.Gender,
		filter.AgeGroup,
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

func (r *DBRepository) SportClubSearch(ctx context.Context, query string, filter dto.SportClubFilters) ([]record.SportClub, error) {
	// TODO Реализовать метод для поиска СО
	return nil, nil
}

func (r *DBRepository) CitySearch(ctx context.Context, query string, filter dto.CityFilters) ([]record.City, error) {
	// TODO Реализовать метод для поиска города
	return nil, nil
}
