package news

type Post struct {
	Id          int    `json:"-" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}
