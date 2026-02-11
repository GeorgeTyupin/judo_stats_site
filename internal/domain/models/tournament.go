package models

type Tournament struct {
	ID        int32  `db:"id"`
	Name      string `db:"name"`
	Type      string `db:"type"`
	Place     string `db:"place"`
	CityID    *int32 `db:"city_id"`
	CountryID *int32 `db:"country_id"`
	Date      string `db:"date"`
	Year      int16  `db:"year"`
	Month     int16  `db:"month"`
	Gender    string `db:"gender"`
	CityLast  string
}
