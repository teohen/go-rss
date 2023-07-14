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

func (apiCfg *apiConfig) handlerGetNewestPosts(w http.ResponseWriter, r *http.Request, user UsersDBModel) {
	var posts []PostsDBModel

	feedFollows, err := apiCfg.dbFeedFollow.getByIdUser(user.UUID, apiCfg.dbOperator)

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error findind newests posts: %s ", err))
		return
	}

	if len(feedFollows) == 0 {
		respondWithError(w, 400, fmt.Sprintf("User does not follow any feeds: %s ", err))
		return
	}

	for _, ff := range feedFollows {
		postForFeed, err := apiCfg.dbPosts.getByFeedId(ff.FeedId, apiCfg.dbOperator)

		if err != nil {
			respondWithError(w, 500, fmt.Sprintf("Error findind newests posts: %s ", err))
			return
		}

		posts = append(posts, postForFeed...)

	}

	respondWithJSON(w, 200, databasesPostsToPosts(posts))
}
