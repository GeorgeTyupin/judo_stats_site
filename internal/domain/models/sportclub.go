package models

type SportClub struct {
	Name           string `db:"name"`
	FullName       string `db:"full_name"`
	Founded        string `db:"founded"`
	Region         string `db:"region"`
	HeadCoach      string `db:"head_coach"`
	Description    string `db:"description"`
	ID             int32  `db:"id"`
	CityID         int32  `db:"city_id"`
	Athletes       int
	ActiveAthletes int
	Coaches        int
	GoldMedals     int
	SilverMedals   int
	BronzeMedals   int
	TotalMedals    int
}
