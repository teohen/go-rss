package main

type User struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Name      string `json:"name"`
	ApiKey    string `json:"api_key"`
}

func databaseUserToUser(dbUser UsersDBModel) User {
	return User{
		ID:        dbUser.UUID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		ApiKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	URL       string `json:"url"`
	UserId    string `json:"user_id"`
}

func databaseFeedToFeed(dbFeed FeedsDBModel) Feed {
	return Feed{
		ID:        dbFeed.Id,
		Name:      dbFeed.Name,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		URL:       dbFeed.URL,
		UserId:    dbFeed.UserId,
	}
}

func databaseFeedsToFeeds(dbFeeds []FeedsDBModel) []Feed {
	feeds := []Feed{}

	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}

type FeedFollow struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	UserId    string `json:"user_id"`
	FeedId    string `json:"feed_id"`
}

func databaseFeedFollowToFeedFollow(dbFeedFollow FeedFollowsDBModel) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.Id,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserId:    dbFeedFollow.UserId,
		FeedId:    dbFeedFollow.FeedId,
	}
}

func databaseFeedFollowsToFeedFollows(dbFeedFollows []FeedFollowsDBModel) []FeedFollow {
	feedFollows := []FeedFollow{}
	for _, dbFeedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, databaseFeedFollowToFeedFollow(dbFeedFollow))
	}
	return feedFollows
}
