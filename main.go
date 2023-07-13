package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	dbUser       *DBUser
	dbOperator   *DB
	dbFeed       *DBFeed
	dbFeedFollow *DBFeedFollows
	dbPosts      *DBPosts
}

type Bar struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Language    string `json:"language"`
}

func main() {

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not in the environment")
	}

	dbOperator := &DB{}
	dbUser := &DBUser{}
	dbFeed := &DBFeed{}
	dbFeedFollows := &DBFeedFollows{}
	dbPosts := &DBPosts{}

	apiCfg := apiConfig{
		dbUser:       dbUser,
		dbOperator:   dbOperator,
		dbFeed:       dbFeed,
		dbFeedFollow: dbFeedFollows,
		dbPosts:      dbPosts,
	}

	go startScraping(apiCfg, 10, time.Minute)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIOS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/health-status", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetAllFeeds)
	v1Router.Post("/feed-follow", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1Router.Get("/feed-follow", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollow))
	v1Router.Delete("/feed-follow/{id}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server running on port: %s", portString)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
