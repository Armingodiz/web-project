package store

import (
	"encoding/json"
	"web-project/db"
	"web-project/models"
)

func NewStore(db *db.DB) Store {
	return &PostgresStore{db}
}

type Store interface {
	// GetById(context context.Context, id int) (pipeline models.PipelineSummery, err error)
	// DeleteExecution(context context.Context, id int) (err error)
	CreateUser(user models.User) error
	GetUser(userName string) (models.User, error)
}

type PostgresStore struct {
	db *db.DB
}

func (s *PostgresStore) CreateUser(user models.User) error {
	_, err := s.db.Connection.Exec("INSERT INTO users (user_name, password, urls) VALUES ($1, $2, $3)", user.Username, user.Password, nil)
	return err
}

func (s *PostgresStore) GetUser(userName string) (models.User, error) {
	var user models.User
	var body interface{}
	err := s.db.Connection.QueryRow("SELECT user_name, password, urls FROM users WHERE user_name = $1", userName).Scan(&user.Username, &user.Password, &body)
	if err != nil {
		return user, err
	}
	if body != nil {
		var urls map[string]string
		json.Unmarshal(body.([]byte), &urls)
		user.Urls = urls
	}
	return user, err
}
