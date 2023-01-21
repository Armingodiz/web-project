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
	CreateUrl(url models.Url) error
	GetUrls(userName string) ([]models.Url, error)
	GetUrl(userName string, address string) (models.Url, error)
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

func (s *PostgresStore) CreateUrl(url models.Url) error {
	_, err := s.db.Connection.Exec("Insert into urls (user_name, address, treshold, failed_times, requests) values ($1, $2, $3, $4, $5)", url.UserName, url.Address, url.Treshold, url.FailedTimes, nil)
	return err
}

func (s *PostgresStore) GetUrls(userName string) ([]models.Url, error) {
	var urls []models.Url
	rows, err := s.db.Connection.Query("SELECT id, user_name, address, treshold, failed_times, requests FROM urls WHERE user_name = $1", userName)
	if err != nil {
		return urls, err
	}
	for rows.Next() {
		var url models.Url
		var body interface{}
		err = rows.Scan(&url.Id, &url.UserName, &url.Address, &url.Treshold, &url.FailedTimes, &body)
		if err != nil {
			return urls, err
		}
		if body != nil {
			var requests []models.Request
			json.Unmarshal(body.([]byte), &requests)
			url.Requests = requests
		}
		urls = append(urls, url)
	}
	return urls, err
}

func (s *PostgresStore) GetUrl(userName string, address string) (models.Url, error) {
	var url models.Url
	var body interface{}
	err := s.db.Connection.QueryRow("SELECT id, user_name, address, threshold, failed_times, requests FROM urls WHERE user_name = $1 AND address = $2", userName, address).Scan(&url.Id, &url.UserName, &url.Address, &url.Treshold, &url.FailedTimes, &body)
	if err != nil {
		return url, err
	}
	if body != nil {
		var requests []models.Request
		json.Unmarshal(body.([]byte), &requests)
		url.Requests = requests
	}
	return url, err
}
