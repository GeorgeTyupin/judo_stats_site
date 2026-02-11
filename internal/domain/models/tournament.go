package models

type Tournament struct {
	Name      string `db:"name"`
	Type      string `db:"type"`
	Place     string `db:"place"`
	Date      string `db:"date"`
	Gender    string `db:"gender"`
	CityID    *int32 `db:"city_id"`
	CountryID *int32 `db:"country_id"`
	ID        int32  `db:"id"`
	Year      int16  `db:"year"`
	Month     int16  `db:"month"`
	CityLast  string
}
