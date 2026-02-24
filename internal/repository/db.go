package repository

import (
	"context"
	"fmt"
	"judo_stats_site/internal/config"
	"judo_stats_site/internal/handlers/dto"
	"judo_stats_site/internal/repository/entity"
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

func (r *DBRepository) GeneralSearch(ctx context.Context, query string) ([]any, error) {
	// TODO Реализовать метод для общего поиска

	return nil, nil
}

func (r *DBRepository) JudokaSearch(ctx context.Context, query string, filter dto.JudokaFilters) ([]entity.Judoka, error) {
	// TODO: Реализовать настоящий поиск в БД
	// Временные моковые данные для тестирования

	mockJudokas := []entity.Judoka{
		{
			ID:               1,
			Name:             "Teddy RINER",
			NAME_RUS:         "Тедди РИНЕР",
			Country:          "France",
			WeightCategories: []string{"+100 кг"},
			BirthDate:        "1989-04-07",
			BirthPlace:       "Paris, France",
			GoldMedals:       15,
			SilverMedals:     3,
			BronzeMedals:     2,
		},
		{
			ID:               2,
			Name:             "Idalys ORTIZ",
			NAME_RUS:         "Идалис ОРТИС",
			Country:          "Cuba",
			WeightCategories: []string{"+78 кг"},
			BirthDate:        "1989-09-27",
			BirthPlace:       "Havana, Cuba",
			GoldMedals:       8,
			SilverMedals:     5,
			BronzeMedals:     4,
		},
		{
			ID:               3,
			Name:             "Ono SHOHEI",
			NAME_RUS:         "ОНО Сёхэй",
			Country:          "Japan",
			WeightCategories: []string{"-73 кг"},
			BirthDate:        "1992-02-03",
			BirthPlace:       "Yamaguchi, Japan",
			GoldMedals:       12,
			SilverMedals:     2,
			BronzeMedals:     1,
		},
		{
			ID:               4,
			Name:             "Clarisse AGBEGNENOU",
			NAME_RUS:         "Кларисс АГБЕНЬЕНУ",
			Country:          "France",
			WeightCategories: []string{"-63 кг"},
			BirthDate:        "1992-10-25",
			BirthPlace:       "Rennes, France",
			GoldMedals:       10,
			SilverMedals:     4,
			BronzeMedals:     3,
		},
		{
			ID:               5,
			Name:             "Lukas KRPALEK",
			NAME_RUS:         "Лукаш КРПАЛЕК",
			Country:          "Czech Republic",
			WeightCategories: []string{"+100 кг", "-100 кг"},
			BirthDate:        "1990-11-15",
			BirthPlace:       "Jihlava, Czech Republic",
			GoldMedals:       7,
			SilverMedals:     6,
			BronzeMedals:     5,
		},
	}

	return mockJudokas, nil
}

func (r *DBRepository) TournamentSearch(ctx context.Context, query string, filter dto.TournamentFilters) ([]entity.Tournament, error) {
	sqlQuery := `
		SELECT tournaments.* FROM tournaments
		LEFT JOIN cities ON tournaments.city_id = cities.id
		LEFT JOIN countries ON tournaments.country_id = countries.id
		LEFT JOIN ussr_republics ON tournaments.republic_id = ussr_republics.id
		WHERE tournaments.name LIKE '%' || $1 || '%'
		AND ($2 = '' OR tournaments.type = $2)
		AND ($3 = '' OR tournaments.gender = $3)
		AND ($4 = 0 OR tournaments.year = $4)
		AND ($5 = '' OR cities.name = $5)
		AND ($6 = '' OR countries.name = $6)
		AND ($7 = '' OR ussr_republics.name = $7);`

	rows, err := r.db.Query(ctx, sqlQuery,
		query,
		filter.Type,
		filter.Gender,
		filter.Year,
		filter.City,
		filter.Country,
		filter.Republic,
	)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса в бд: %w", err)
	}

	tournaments, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.Tournament])
	if err != nil {
		return nil, fmt.Errorf("ошибка преобразования rows к entity.Tournament: %w", err)
	}

	return tournaments, nil
}

func (r *DBRepository) SportClubSearch(ctx context.Context, query string, filter dto.SportClubFilters) ([]entity.SportClub, error) {
	// TODO Реализовать метод для поиска СО

	return nil, nil
}

func (r *DBRepository) CitySearch(ctx context.Context, query string, filter dto.CityFilters) ([]entity.City, error) {
	// TODO Реализовать метод для поиска города

	return nil, nil
}
