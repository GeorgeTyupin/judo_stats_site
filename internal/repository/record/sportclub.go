package record

type SportClub struct {
	Name        string `db:"name"`
	FullName    string `db:"full_name"`
	Founded     string `db:"founded"`
	Region      string `db:"region"`
	HeadCoach   string `db:"head_coach"`
	Description string `db:"description"`
	ID          int64  `db:"id"`
	CityID      int64  `db:"city_id"`
}
