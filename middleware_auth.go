package main

import (
	"fmt"
	"net/http"

	"github.com/teohen/go-rss/auth"
)

type authHandler func(http.ResponseWriter, *http.Request, UsersDBModel)

func (apicfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := apicfg.dbUser.getBy(apiKey, apicfg.dbOperator)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldnt get user: %s", err))
			return
		}

		handler(w, r, user)
	}
}
