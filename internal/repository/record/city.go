package record

type City struct {
	ID          int32  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	RepublicID  *int64 `db:"republic_id"`
}
