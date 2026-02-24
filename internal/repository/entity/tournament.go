package entity

import "time"

type Tournament struct {
	Name       string    `db:"name"`
	Type       string    `db:"type"`
	Place      string    `db:"place"`
	Date       string    `db:"date"`
	Gender     string    `db:"gender"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	ID         int64     `db:"id"`
	CityID     *int32    `db:"city_id"`
	CountryID  *int32    `db:"country_id"`
	RepublicID *int32    `db:"republic_id"`
	Year       int16     `db:"year"`
	Month      int16     `db:"month"`
}
