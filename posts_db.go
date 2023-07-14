package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type DBPosts struct {
	collection []PostsDBModel
}

type PostsDBModel struct {
	Id          string
	CreatedAt   string
	UpdatedAt   string
	Title       string
	Description string
	PublishedAt string
	URL         string
	FeedId      string
}

func (dbPost *DBPosts) save(post PostsDBModel, dbOperator *DB) (PostsDBModel, error) {
	if !dbOperator.isConnected() {
		dbOperator.connect()
	}

	post.Id = uuid.NewString()

	msg, err := dbOperator.execute(post.Id, "save")

	if err != nil {
		fmt.Println(fmt.Printf("ERROR: %s", err))
		return post, errors.New(fmt.Sprintf("Error trying to save post: %s", post.Id))
	}

	dbPost.collection = append(dbPost.collection, post)

	fmt.Println(fmt.Sprintf("SUCCESS: %s", msg))
	return post, nil
}

func (dbPost *DBPosts) getAll(dbOperator *DB) ([]PostsDBModel, error) {
	var user []PostsDBModel

	if !dbOperator.isConnected() {
		dbOperator.connect()
	}

	msg, err := dbOperator.execute("nil", "get all")

	if err != nil {
		fmt.Println(fmt.Printf("ERROR: %s", err))
		return user, errors.New(fmt.Sprintf("error trying to get all posts %s", msg))
	}

	return dbPost.collection, nil
}

func (dbPost *DBPosts) update(postToUpdate PostsDBModel, dbOperator *DB) (PostsDBModel, error) {
	var updatedPost PostsDBModel

	if !dbOperator.isConnected() {
		dbOperator.connect()
	}

	msg, err := dbOperator.execute(postToUpdate.URL, "update")

	if err != nil {
		fmt.Println(fmt.Printf("ERROR: %s", err))
		return updatedPost, errors.New(fmt.Sprintf("error trying to update post %s", msg))
	}

	for i, post := range dbPost.collection {
		if post.URL == postToUpdate.URL {
			dbPost.collection[i].Id = post.Id
			dbPost.collection[i].Title = postToUpdate.Title
			dbPost.collection[i].Description = postToUpdate.Description
			dbPost.collection[i].CreatedAt = post.CreatedAt
			dbPost.collection[i].UpdatedAt = time.Now().String()
			dbPost.collection[i].PublishedAt = postToUpdate.CreatedAt
			dbPost.collection[i].FeedId = post.FeedId
			return dbPost.collection[i], nil
		}
	}

	return updatedPost, nil
}

func (dbPost *DBPosts) getByURL(url string, dbOperator *DB) (PostsDBModel, error) {
	var post PostsDBModel

	if !dbOperator.isConnected() {
		dbOperator.connect()
	}

	msg, err := dbOperator.execute(url, "get by url")

	if err != nil {
		fmt.Println(fmt.Printf("ERROR: %s", err))
		return post, errors.New(fmt.Sprintf("error trying to get post by url %s", msg))
	}

	for _, post := range dbPost.collection {
		if post.URL == url {
			return post, nil
		}
	}

	return post, nil
}
