package main

import (
	"log"
	"sync"
	"time"
)

func startScraping(apiCfg apiConfig, concurrency int, interval time.Duration) {
	log.Printf("Scrapping on %v goroutines on the interval of: %s", concurrency, interval)

	ticker := time.NewTicker(interval)

	for ; ; <-ticker.C {
		feeds, err := apiCfg.dbFeed.getNextFeedsToFetch(concurrency, apiCfg.dbOperator)

		if err != nil {
			log.Println("error fetching feeds", err)
			continue
		}

		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(wg, apiCfg, feed)
		}

		wg.Wait()
	}
}

func scrapeFeed(wg *sync.WaitGroup, apiCfg apiConfig, feed FeedsDBModel) {
	defer wg.Done()

	_, err := apiCfg.dbFeed.markFeedAsFetched(feed.Id, apiCfg.dbOperator)

	if err != nil {
		log.Println("error updating feed", err)
		return
	}

	rssFeed, err := urlToFeed(feed.URL)
	if err != nil {
		log.Println("error fetching feed", err)
	}

	for _, post := range rssFeed {
		dbPost, err := apiCfg.dbPosts.getByURL(post.Link, apiCfg.dbOperator)

		if err != nil {
			log.Println("error on getting post", err)
		}

		if dbPost.URL == "" {
			log.Println("saving post")
			apiCfg.dbPosts.save(PostsDBModel{
				Title:       post.Link,
				Description: post.Description,
				URL:         post.Link,
				CreatedAt:   time.Now().String(),
				UpdatedAt:   time.Now().String(),
				FeedId:      dbPost.FeedId,
				PublishedAt: post.CreatedAt,
			}, apiCfg.dbOperator)
		} else {
			log.Println("updating post")
			dbPost.Title = post.Link
			dbPost.Description = post.Description
			apiCfg.dbPosts.update(dbPost, apiCfg.dbOperator)
		}
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed))
}
