package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user UsersDBModel) {
	type Parameters struct {
		FeedId string `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := Parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s ", err))
		return
	}

	feedFollow, err := apiCfg.dbFeedFollow.save(FeedFollowsDBModel{
		Id:        uuid.NewString(),
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
		UserId:    user.UUID,
		FeedId:    params.FeedId,
	}, apiCfg.dbOperator)

	if err != nil {
		statusCode := 400
		if err.Error() == "Constraint error: id user + id feed" {
			statusCode = 409
		}

		respondWithError(w, statusCode, fmt.Sprintf("Error creatig feed follow"))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}
