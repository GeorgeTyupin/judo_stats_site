package record

type USSRRepublic struct {
	ID          int32  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}
