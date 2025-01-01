package news

type User struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
