package repository

import (
	"context"
	"fmt"
	"judo_stats_site/internal/config"
	"judo_stats_site/internal/domain/dto"
	"judo_stats_site/internal/domain/models"
	"log/slog"
	"time"

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

func (r *DBRepository) GeneralSearch(query string) []any {
	// TODO Реализовать метод для общего поиска

	return nil
}

func (r *DBRepository) JudokaSearch(query string, filter dto.JudokaFilters) []models.Judoka {
	// TODO: Реализовать настоящий поиск в БД
	// Временные моковые данные для тестирования

	mockJudokas := []models.Judoka{
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

	return mockJudokas
}

func (r *DBRepository) TournamentSearch(query string, filter dto.TournamentFilters) []models.Tournament {
	// TODO Реализовать метод для поиска турнира

	return nil
}

func (r *DBRepository) SportClubSearch(query string, filter dto.SportClubFilters) []models.SportClub {
	// TODO Реализовать метод для поиска СО

	return nil
}

func (r *DBRepository) CitySearch(query string, filter dto.CityFilters) []models.City {
	// TODO Реализовать метод для поиска города

	return nil
}
