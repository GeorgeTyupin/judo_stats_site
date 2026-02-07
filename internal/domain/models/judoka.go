package models

type Judoka struct {
	ID               int64    `db:"id"`
	Name             string   `db:"name"`
	NAME_RUS         string   `db:"name_rus"`
	Country          string   `db:"country"`
	WeightCategories []string `db:"weight_category"`
	BirthDate        string   `db:"birth_date"`
	BirthPlace       string   `db:"birth_place"`
	GoldMedals       uint8
	SilverMedals     uint8
	BronzeMedals     uint8
}
