package record

type Judoka struct {
	ID               int64    `db:"id"`
	LastName         string   `db:"last_name"`
	FirstName        string   `db:"first_name"`
	LastNameRus      string   `db:"last_name_rus"`
	FirstNameRus     string   `db:"first_name_rus"`
	Gender           string   `db:"gender"`
	Country          string   `db:"country"`
	City             string   `db:"city"`
	SportClub        string   `db:"sport_club"`
	WeightCategories []string `db:"weight_category"`
	BirthDate        string   `db:"birth_date"`
	BirthPlace       string   `db:"birth_place"`
}
