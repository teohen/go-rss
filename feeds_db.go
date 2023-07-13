package main

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

type DBFeed struct {
	collection []FeedsDBModel
}

type FeedsDBModel struct {
	Id            string
	CreatedAt     string
	UpdatedAt     string
	Name          string
	URL           string
	UserId        string
	LastFetchedAt string
}

func (dbFeed *DBFeed) save(feed FeedsDBModel, dbOperator *DB) (FeedsDBModel, error) {
	if !dbOperator.isConnected() {
		dbOperator.connect()
	}

	msg, err := dbOperator.execute(feed.Id, "save")

	if err != nil {
		fmt.Println(fmt.Printf("ERROR: %s", err))
		return feed, errors.New(fmt.Sprintf("Error trying to save feed: %s", feed.Id))
	}

	dbFeed.collection = append(dbFeed.collection, feed)

	fmt.Println(fmt.Sprintf("SUCCESS: %s", msg))
	return feed, nil
}

func (dbFeed *DBFeed) getAll(dbOperator *DB) ([]FeedsDBModel, error) {
	var user []FeedsDBModel

	if !dbOperator.isConnected() {
		dbOperator.connect()
	}

	msg, err := dbOperator.execute("nil", "get all")

	if err != nil {
		fmt.Println(fmt.Printf("ERROR: %s", err))
		return user, errors.New(fmt.Sprintf("error trying to get all feeds %s", msg))
	}

	return dbFeed.collection, nil
}

func (dbFeed *DBFeed) getNextFeedsToFetch(limit int, dbOperator *DB) ([]FeedsDBModel, error) {
	feedsToFetch := []FeedsDBModel{}

	allFeeds, err := dbFeed.getAll(dbOperator)

	if err != nil {
		fmt.Println(fmt.Printf("ERROR: %s", err))
		return feedsToFetch, errors.New("error trying to get feeds to fetch")
	}

	if len(allFeeds) == 0 {
		return feedsToFetch, nil
	}

	if len(allFeeds) < limit {
		limit = len(allFeeds)
	}

	sort.Slice(allFeeds, func(i, j int) bool {
		if allFeeds[i].LastFetchedAt == "" {
			return false
		}
		if allFeeds[j].LastFetchedAt == "" {
			return true
		}
		return allFeeds[i].LastFetchedAt > allFeeds[j].LastFetchedAt
	})

	return allFeeds[:limit], nil
}

func (dbFeed *DBFeed) markFeedAsFetched(idFeed string, dbOperator *DB) (bool, error) {

	if !dbOperator.isConnected() {
		dbOperator.connect()
	}

	msg, err := dbOperator.execute(idFeed, "update feed")

	if err != nil {
		fmt.Println(fmt.Printf("ERROR: %s", err))
		return false, errors.New("error trying to update feeds")
	}

	fmt.Println(fmt.Sprintf("SUCCESS: %s", msg))

	for i, feed := range dbFeed.collection {
		if feed.Id == idFeed {
			dbFeed.collection[i].LastFetchedAt = time.Now().String()
			return true, nil
		}
	}

	return false, errors.New("feed not found on db")
}
