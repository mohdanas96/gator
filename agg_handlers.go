package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mohdanas96/gator/internal/api"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("duration is required as a parameter, eg :- '1s', '1m', '1h'")
	}

	duration, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("something went wrong while parsing duration :: %v", err)
	}

	ticker := time.NewTicker(duration)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("something went wrong while fetching next feed to fetch :: %v", err)
	}

	log.Println("found a feed to fetch")

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("something went wrong while updating feed :: %v", err)
	}

	rssFeed, err := api.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("something went wrong while fetching rss feed :: %v", err)
	}

	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))

	return nil
}
