package models

type SportClub struct {
	ID             int32  `db:"id"`
	Name           string `db:"name"`
	FullName       string `db:"full_name"`
	Founded        string `db:"founded"`
	CityID         int32  `db:"city_id"`
	Region         string `db:"region"`
	HeadCoach      string `db:"head_coach"`
	Description    string `db:"description"`
	Athletes       int
	ActiveAthletes int
	Coaches        int
	GoldMedals     int
	SilverMedals   int
	BronzeMedals   int
	TotalMedals    int
}
