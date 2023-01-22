package requester

import (
	"time"
	"web-project/models"
	"web-project/store"
	"web-project/utils"
)

func GetRequester(interval int, storage store.Store) (*Requester, error) {
	if cronRequester == nil {
		cronRequester = &Requester{
			Interval: interval,
			Store:    storage,
		}
		urls, err := cronRequester.Store.GetAllUrls()
		if err != nil {
			return nil, err
		}
		cronRequester.Urls = urls
	}
	go cronRequester.StartMonitoring()
	return cronRequester, nil
}

type Requester struct {
	Interval int
	Urls     []models.Url
	Store    store.Store
}

var cronRequester *Requester

func (r *Requester) AddUrl(url models.Url) {
	r.Urls = append(r.Urls, url)
}

func (r *Requester) StartMonitoring() {
	for {
		for _, url := range r.Urls {
			go r.monitorUrl(url)
		}
		time.Sleep(time.Duration(r.Interval) * time.Second)
	}
}

func (r *Requester) monitorUrl(url models.Url) {
	res, _ := utils.SendRequest(url)
	r.Store.AddRequest(url.Id, res.Result)
}
