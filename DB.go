package main

import (
	"errors"
	"fmt"
)

type DB struct {
	connected bool
}

func (dbOperator *DB) connect() {
	if !dbOperator.connected {
		dbOperator.connected = true
	}
	fmt.Println("Connected to DB")
}

func (dbOperator *DB) isConnected() bool {
	return dbOperator.connected
}

func (dbOperator *DB) disconnect() {
	if dbOperator.connected == true {
		dbOperator.connected = false
	}
	fmt.Println("Disconnected from DB")
}

func (dbOperator *DB) execute(identifier string, operation string) (string, error) {
	defer dbOperator.disconnect()

	if !dbOperator.isConnected() {
		return "", errors.New("Not connected")
	}
	msg := fmt.Sprintf("Operation on DB: %s with values: %s", operation, identifier)
	return msg, nil
}
