package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type RSSFeed struct {
	Title       string
	Link        string
	Description string
	Language    string
}

func urlToFeed(url string) ([]RSSFeed, error) {

	rssFeed := []RSSFeed{}

	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)

	if err != nil {
		return rssFeed, err
	}

	defer resp.Body.Close()

	body, errIO := ioutil.ReadAll(resp.Body)

	if errIO != nil {
		return rssFeed, errIO
	}

	err = json.Unmarshal([]byte(string(body)), &rssFeed)

	return rssFeed, nil
}
