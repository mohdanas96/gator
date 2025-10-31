package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mohdanas96/gator/internal/api"
	"github.com/mohdanas96/gator/internal/database"
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

	for _, v := range rssFeed.Channel.Item {
		parsedDate, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", v.PubDate)
		if err != nil {
			return fmt.Errorf("something went wrong while parsing date :: %v", err)
		}

		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       v.Title,
			Url:         v.Link,
			Description: sql.NullString{String: v.Description, Valid: true},
			PublishedAt: parsedDate,
			FeedID:      feed.ID,
		}
		_, err = s.db.CreatePost(context.Background(), params)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}

	return nil
}
