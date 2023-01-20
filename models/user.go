package models

type User struct {
	Username string            `db:"user_name" json:"user_name"`
	Password string            `db:"password" json:"password"`
	Urls     map[string]string `db:"urls" json:"urls"`
}
