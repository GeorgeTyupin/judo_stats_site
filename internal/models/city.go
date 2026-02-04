package models

type CityData struct {
	ID             int64  `db:"id"`
	Name           string `db:"name"`
	Region         string `db:"region"`
	Population     int    `db:"population"`
	Founded        string `db:"founded"`
	Description    string `db:"description"`
	SportClubs     int
	Athletes       int
	ActiveAthletes int
	Coaches        int
	GoldMedals     int
	SilverMedals   int
	BronzeMedals   int
	TotalMedals    int
}
