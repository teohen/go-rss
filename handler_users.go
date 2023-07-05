package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/teohen/go-rss/auth"
)

type UsersDBModel struct {
	UUID      string
	Name      string
	CreatedAt string
	UpdatedAt string
	ApiKey    string
}

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type Parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := Parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s ", err))
		return
	}

	user, err := apiCfg.dbOperator.save(UsersDBModel{
		Name:      params.Name,
		UUID:      uuid.NewString(),
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creatig user"))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetApiKey(r.Header)

	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Unauthorized"))
		return
	}

	user, err := apiCfg.dbOperator.getBy(apiKey)

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Internal Server Error"))
		return
	}

	if user.Name == "" {
		respondWithError(w, 400, fmt.Sprintf("Bad Request"))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}
