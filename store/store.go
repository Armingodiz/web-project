package store

import (
	"encoding/json"
	"time"
	"web-project/db"
	"web-project/models"
)

func NewStore(db *db.DB) Store {
	return &PostgresStore{db}
}

type Store interface {
	CreateUser(user models.User) error
	GetUser(userName string) (models.User, error)
	CreateUrl(url models.Url) (string, error)
	GetUrls(userName string) ([]models.Url, error)
	GetUrl(urlId string) (models.Url, error)
	AddRequest(urlId string, result int) error
	GetRequests(urlId string) (requests []models.Request, err error)
	GetAllUrls() ([]models.Url, error)
	GetAlerts(urlId string) ([]models.Alert, error)
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

func (s *PostgresStore) CreateUrl(url models.Url) (string, error) {
	_, err := s.db.Connection.Exec("Insert into urls (user_name, address, treshold, failed_times, requests) values ($1, $2, $3, $4, $5)", url.UserName, url.Address, url.Treshold, url.FailedTimes, nil)
	if err != nil {
		return "", err
	}
	var id string
	err = s.db.Connection.QueryRow("SELECT id FROM urls WHERE user_name = $1 AND address = $2", url.UserName, url.Address).Scan(&id)
	return id, err
}

func (s *PostgresStore) GetUrls(userName string) ([]models.Url, error) {
	var urls []models.Url
	rows, err := s.db.Connection.Query("SELECT id, user_name, address, treshold, failed_times FROM urls WHERE user_name = $1", userName)
	if err != nil {
		return urls, err
	}
	for rows.Next() {
		var url models.Url
		err = rows.Scan(&url.Id, &url.UserName, &url.Address, &url.Treshold, &url.FailedTimes)
		if err != nil {
			return urls, err
		}
		requests, err := s.GetRequests(url.Id)
		if err == nil {
			url.Requests = requests
		}
		urls = append(urls, url)
	}
	return urls, err
}

func (s *PostgresStore) GetUrl(urlId string) (models.Url, error) {
	var url models.Url
	err := s.db.Connection.QueryRow("SELECT id, user_name, address, treshold, failed_times FROM urls WHERE id = $1", urlId).Scan(&url.Id, &url.UserName, &url.Address, &url.Treshold, &url.FailedTimes)
	if err != nil {
		return url, err
	}
	requests, err := s.GetRequests(url.Id)
	if err == nil {
		url.Requests = requests
	}
	return url, err
}

func (s *PostgresStore) AddRequest(urlId string, result int) (err error) {
	_, err = s.db.Connection.Exec("INSERT into url_requests (url_id, result) values ($1, $2)", urlId, result)
	if err != nil {
		return
	}
	if int(result/100) != 2 {
		_, err = s.db.Connection.Exec("UPDATE urls SET failed_times = failed_times + 1 WHERE id = $1", urlId)
		if err != nil {
			return
		}
		var currentFailedTimes, treshold int
		err = s.db.Connection.QueryRow("SELECT failed_times, treshold FROM urls WHERE id = $1", urlId).Scan(&currentFailedTimes, &treshold)
		if err != nil {
			return
		}
		if currentFailedTimes > treshold {
			_, err = s.db.Connection.Exec("INSERT INTO alerts (url_id, message) values ($1, $2)", urlId, "Your url address failed times passed treshold at "+time.Now().GoString())
		}
	}
	return err
}

func (s *PostgresStore) GetRequests(urlId string) (requests []models.Request, err error) {
	requests = []models.Request{}
	rows, err := s.db.Connection.Query("SELECT url_id, result from url_requests where url_id = $1", urlId)
	if err != nil {
		return
	}
	for rows.Next() {
		var req models.Request
		err = rows.Scan(&req.UrlId, &req.Result)
		if err != nil {
			return
		}
		requests = append(requests, req)
	}
	return
}

func (s *PostgresStore) GetAllUrls() ([]models.Url, error) {
	var urls []models.Url
	rows, err := s.db.Connection.Query("SELECT id, user_name, address, treshold, failed_times FROM urls")
	if err != nil {
		return urls, err
	}
	for rows.Next() {
		var url models.Url
		err = rows.Scan(&url.Id, &url.UserName, &url.Address, &url.Treshold, &url.FailedTimes)
		if err != nil {
			return urls, err
		}
		requests, err := s.GetRequests(url.Id)
		if err == nil {
			url.Requests = requests
		}
		urls = append(urls, url)
	}
	return urls, err
}

func (s *PostgresStore) GetAlerts(urlId string) ([]models.Alert, error) {
	var alerts []models.Alert
	rows, err := s.db.Connection.Query("SELECT url_id, message FROM alerts WHERE url_id = $1", urlId)
	if err != nil {
		return alerts, err
	}
	for rows.Next() {
		var alert models.Alert
		err = rows.Scan(&alert.UrlId, &alert.Message)
		if err != nil {
			return alerts, err
		}
		alerts = append(alerts, alert)
	}
	return alerts, err
}
