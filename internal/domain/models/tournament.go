package models

type Tournament struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	Type     string `db:"type"`
	Place    string `db:"place"`
	CityID   *int64 `db:"city_id"`
	Date     string `db:"date"`
	Year     int16  `db:"year"`
	Month    int8   `db:"month"`
	Gender   string `db:"gender"`
	CityLast string
}
