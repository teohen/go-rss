package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user UsersDBModel) {
	type Parameters struct {
		Name string `json:"name"`
		URL  string `json:"URL"`
	}
	decoder := json.NewDecoder(r.Body)

	params := Parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s ", err))
		return
	}

	feed, err := apiCfg.dbFeed.save(FeedsDBModel{
		Id:        uuid.NewString(),
		Name:      params.Name,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
		URL:       params.URL,
		UserId:    user.UUID,
	}, apiCfg.dbOperator)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creatig feed"))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.dbFeed.getAll(apiCfg.dbOperator)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error getting all feeds"))
		return
	}

	respondWithJSON(w, 200, databaseFeedsToFeeds(feeds))
}
