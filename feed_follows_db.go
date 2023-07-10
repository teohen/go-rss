package main

import (
	"errors"
	"fmt"
)

type DBFeedFollows struct {
	collection []FeedFollowsDBModel
}

type FeedFollowsDBModel struct {
	Id        string
	UserId    string
	FeedId    string
	CreatedAt string
	UpdatedAt string
}

func (dbFeedFollow *DBFeedFollows) save(feedFollow FeedFollowsDBModel, dbOperator *DB) (FeedFollowsDBModel, error) {
	if !dbOperator.isConnected() {
		dbOperator.connect()
	}

	existentFfModel, err := dbFeedFollow.getByIdUserAndIdFeed(feedFollow.UserId, feedFollow.FeedId, dbOperator)

	if err != nil {
		fmt.Println(fmt.Printf("ERROR on save: %s", err))
		return feedFollow, errors.New(fmt.Sprintf("Error trying to save feed follow: %s", feedFollow.Id))
	}

	if existentFfModel.Id != "" {
		fmt.Println("Constraint error: id user + id feed")
		return existentFfModel, errors.New(fmt.Sprintf("Constraint error: id user + id feed"))
	}

	if !dbOperator.isConnected() {
		dbOperator.connect()
	}

	msg, err := dbOperator.execute(feedFollow.Id, "save")

	if err != nil {
		fmt.Println(fmt.Printf("ERROR on save 2: %s", err))
		return feedFollow, errors.New(fmt.Sprintf("Error trying to save feed follow: %s", feedFollow.Id))
	}

	dbFeedFollow.collection = append(dbFeedFollow.collection, feedFollow)

	fmt.Println(fmt.Sprintf("SUCCESS: %s", msg))
	return feedFollow, nil
}

func (dbFeedFollow *DBFeedFollows) getByIdUser(idUser string, dbOperator *DB) (FeedFollowsDBModel, error) {
	var feedFollowModel FeedFollowsDBModel

	if !dbOperator.isConnected() {
		dbOperator.connect()
	}

	msg, err := dbOperator.execute(idUser, "get by id user")

	if err != nil {
		fmt.Println(fmt.Printf("ERROR: %s", err))
		return feedFollowModel, errors.New(fmt.Sprintf("error trying to get feed follow: %s", idUser))
	}

	fmt.Println(fmt.Sprintf("SUCCESS %s", msg))

	for _, feedFollow := range dbFeedFollow.collection {
		if feedFollow.UserId == idUser {
			return feedFollow, nil
		}
	}

	return feedFollowModel, nil
}

func (dbFeedFollow *DBFeedFollows) getByIdFeed(idFeed string, dbOperator *DB) (FeedFollowsDBModel, error) {
	var feedFollowModel FeedFollowsDBModel

	if !dbOperator.isConnected() {
		dbOperator.connect()
	}

	msg, err := dbOperator.execute(idFeed, "get by id feed")

	if err != nil {
		fmt.Println(fmt.Printf("ERROR: %s", err))
		return feedFollowModel, errors.New(fmt.Sprintf("error trying to get feed follow: %s", idFeed))
	}

	fmt.Println(fmt.Sprintf("SUCCESS %s", msg))

	for _, feedFollow := range dbFeedFollow.collection {
		if feedFollow.FeedId == idFeed {
			return feedFollow, nil
		}
	}

	return feedFollowModel, nil
}

func (dbFeedFollow *DBFeedFollows) getByIdUserAndIdFeed(idUser string, idFeed string, dbOperator *DB) (FeedFollowsDBModel, error) {
	var feedFollowModel FeedFollowsDBModel

	if !dbOperator.isConnected() {
		dbOperator.connect()
	}

	msg, err := dbOperator.execute(fmt.Sprintf("Id User: %s - Id Feed: %s", idUser, idFeed), "get by id user and id feed")

	if err != nil {
		fmt.Println(fmt.Printf("ERROR on getByUserAndIdFeed: %s", err))
		return feedFollowModel, errors.New("error trying to get feed follow by id user and id feed")
	}

	fmt.Println(fmt.Sprintf("SUCCESS: %s", msg))

	for _, ffM := range dbFeedFollow.collection {
		if ffM.UserId == idUser && ffM.FeedId == idFeed {
			return ffM, nil
		}
	}

	return feedFollowModel, nil
}
