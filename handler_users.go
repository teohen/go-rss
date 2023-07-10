package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

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

	user, err := apiCfg.dbUser.save(UsersDBModel{
		Name:      params.Name,
		UUID:      uuid.NewString(),
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}, apiCfg.dbOperator)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creatig user"))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user UsersDBModel) {
	respondWithJSON(w, 200, databaseUserToUser(user))
}
