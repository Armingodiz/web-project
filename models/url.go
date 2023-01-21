package models

type Url struct {
	Id          string    `json:"id" db:"id"`
	UserName    string    `db:"user_name" json:"user_name"`
	Address     string    `db:"address" json:"address"`
	Treshold    int       `db:"treshold" json:"treshold"`
	FailedTimes int       `db:"failed_times" json:"failed_times"`
	Requests    []Request `db:"requests" json:"requests"`
}
