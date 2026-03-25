package record

type Country struct {
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}
