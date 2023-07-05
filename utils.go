package main

import (
	"crypto/sha256"
	"fmt"
)

func encrypt(data string) string {
	hash := sha256.Sum256([]byte(data))
	hashString := fmt.Sprintf("%x", hash)
	return hashString
}
