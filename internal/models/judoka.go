package models

type Judoka struct {
	ID             int64  `db:"id"`
	Name           string `db:"name"`
	NAME_RUS       string `db:"name_rus"`
	Country        string `db:"country"`
	CountryFlag    string `db:"country_flag"`
	WeightCategory string `db:"weight_category"`
	BirthDate      string `db:"birth_date"`
	BirthPlace     string `db:"birth_place"`
	CityID         *int64 `db:"city_id"`
	SportClubID    *int64 `db:"sport_club_id"`
	SportClub      string `db:"sport_club"`
	Coach          string `db:"coach"`
	GoldMedals     uint8
	SilverMedals   uint8
	BronzeMedals   uint8
}
