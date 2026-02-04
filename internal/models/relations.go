package models

type TournamentResult struct {
	ID             int64  `db:"id"`
	TournamentID   int64  `db:"tournament_id"`
	JudokaID       int64  `db:"judoka_id"`
	WeightCategory string `db:"weight_category"`
	Rank           int    `db:"rank"`
}

type JudokaCityRelation struct {
	ID       int64 `db:"id"`
	JudokaID int64 `db:"judoka_id"`
	CityID   int64 `db:"city_id"`
}

type JudokaSportClubRelation struct {
	ID          int64 `db:"id"`
	JudokaID    int64 `db:"judoka_id"`
	SportClubID int64 `db:"sport_club_id"`
}

type SportClubCityRelation struct {
	ID          int64 `db:"id"`
	SportClubID int64 `db:"sport_club_id"`
	CityID      int64 `db:"city_id"`
}
