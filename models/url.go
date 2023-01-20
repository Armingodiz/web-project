package models

type Url struct {
	Id          uint      `json:"id" db:"id"`
	UserName    string    `db:"user_name" json:"user_name"`
	Address     string    `db:"address" json:"address"`
	Threshold   int       `db:"threshold" json:"threshold"`
	FailedTimes int       `db:"failed_times" json:"failed_times"`
	Requests    []Request `db:"requests" json:"requests"`
}
