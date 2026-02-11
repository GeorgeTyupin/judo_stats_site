package models

type City struct {
	ID          int32  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	RepublicID  *int64 `db:"republic_id"`

	// Computed fields (not stored in DB)
	Region         string
	Population     int
	Founded        string
	SportClubs     int
	Athletes       int
	ActiveAthletes int
	Coaches        int
	GoldMedals     int
	SilverMedals   int
	BronzeMedals   int
	TotalMedals    int
}
