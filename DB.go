package main

import (
	"errors"
	"fmt"
)

type ConnectionDB interface {
	save() (UsersDBModel, error)
}

type DB struct {
	connected bool
	operation string
	users     []UsersDBModel
}

func (dbOperator *DB) execute(identifier string) (string, error) {
	if dbOperator.connected != true {
		return "", errors.New("Not connected")
	}
	msg := fmt.Sprintf("Operation on DB: %s with values: %s", dbOperator.operation, identifier)
	return msg, nil
}

func (dbOperator *DB) save(user UsersDBModel) (UsersDBModel, error) {
	dbOperator.operation = "save"

	user.ApiKey = encrypt(user.Name + user.UUID)

	msg, err := dbOperator.execute(user.Name)
	dbOperator.users = append(dbOperator.users, user)

	if err != nil {

		fmt.Println(fmt.Printf("ERROR: %s", err))
		return user, errors.New(fmt.Sprintf("Error trying to save user: %s", user.Name))
	}

	fmt.Println(fmt.Sprintf("SUCCESS: %s", msg))
	return user, nil
}

func (dbOperator *DB) getBy(apiKey string) (UsersDBModel, error) {
	var user UsersDBModel
	dbOperator.operation = "get by api key"

	msg, err := dbOperator.execute(apiKey)

	if err != nil {
		fmt.Println(fmt.Printf("ERROR: %s", err))
		return user, errors.New("Error trying to get user by api key")
	}

	fmt.Println(fmt.Sprintf("SUCCESS: %s", msg))

	for _, item := range dbOperator.users {
		if item.ApiKey == apiKey {

			user = item
			break
		}
	}

	return user, nil
}
