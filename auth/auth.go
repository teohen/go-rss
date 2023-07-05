package auth

import (
	"errors"
	"log"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	header := headers.Get("Authorization")
	if header == "" {
		err := "header unformatted"
		log.Printf("%s", err)
		return "", errors.New(err)

	}

	authorization := strings.Split(header, " ")

	if len(authorization) != 2 {
		err := "header unformatted"
		log.Printf("%s", err)
		return "", errors.New(err)
	}

	if authorization[0] != "ApiKey" {
		err := "header type not allowed"
		log.Printf("%s", err)
		return "", errors.New(err)
	}

	return authorization[1], nil
}
