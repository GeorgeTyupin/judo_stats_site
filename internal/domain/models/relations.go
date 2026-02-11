package models

type TournamentResult struct {
	ID             int32  `db:"id"`
	TournamentID   int32  `db:"tournament_id"`
	JudokaID       int32  `db:"judoka_id"`
	WeightCategory string `db:"weight_category"`
	Rank           int    `db:"rank"`
}

type JudokaCityRelation struct {
	ID       int32 `db:"id"`
	JudokaID int32 `db:"judoka_id"`
	CityID   int32 `db:"city_id"`
}

type JudokaSportClubRelation struct {
	ID          int32 `db:"id"`
	JudokaID    int32 `db:"judoka_id"`
	SportClubID int32 `db:"sport_club_id"`
}

type SportClubCityRelation struct {
	ID          int32 `db:"id"`
	SportClubID int32 `db:"sport_club_id"`
	CityID      int32 `db:"city_id"`
}

type JudokaCountryRelation struct {
	ID        int32 `db:"id"`
	JudokaID  int64 `db:"judoka_id"`
	CountryID int64 `db:"country_id"`
}

type JudokaRepublicRelation struct {
	ID         int32 `db:"id"`
	JudokaID   int64 `db:"judoka_id"`
	RepublicID int64 `db:"republic_id"`
}
