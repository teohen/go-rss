package main

import (
	"errors"
	"fmt"
)

type DBUser struct {
	collection []UsersDBModel
}

type UsersDBModel struct {
	UUID      string
	Name      string
	CreatedAt string
	UpdatedAt string
	ApiKey    string
}

func (dbUser *DBUser) save(user UsersDBModel, dbOperator *DB) (UsersDBModel, error) {
	if !dbOperator.isConnected() {
		dbOperator.connect()
	}

	user.ApiKey = encrypt(user.Name + user.UUID)

	msg, err := dbOperator.execute(user.Name, "save")

	if err != nil {
		fmt.Println(fmt.Printf("ERROR: %s", err))
		return user, errors.New(fmt.Sprintf("Error trying to save user: %s", user.Name))
	}

	dbUser.collection = append(dbUser.collection, user)

	fmt.Println(fmt.Sprintf("SUCCESS: %s", msg))
	return user, nil
}

func (dbUser *DBUser) getBy(apiKey string, dbOperator *DB) (UsersDBModel, error) {
	var user UsersDBModel

	if !dbOperator.isConnected() {
		dbOperator.connect()
	}

	msg, err := dbOperator.execute(apiKey, "get by api key")

	if err != nil {
		fmt.Println(fmt.Printf("ERROR: %s", err))
		return user, errors.New(fmt.Sprintf("error trying to get user by api key: %s", apiKey))
	}

	for _, user := range dbUser.collection {
		if user.ApiKey == apiKey {
			fmt.Println(fmt.Sprintf("SUCCESS: %s", msg))
			return user, nil
		}
	}

	return user, errors.New(fmt.Sprintf("User %s not found on DB", apiKey))
}
