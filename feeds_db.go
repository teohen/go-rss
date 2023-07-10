package main

import (
	"errors"
	"fmt"
)

type DBFeed struct {
	collection []FeedsDBModel
}

type FeedsDBModel struct {
	Id        string
	CreatedAt string
	UpdatedAt string
	Name      string
	URL       string
	UserId    string
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
