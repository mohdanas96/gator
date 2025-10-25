package main

import (
	"context"
	"errors"
	"fmt"
	"html"

	"github.com/mohdanas96/gator/internal/api"
)

func aggHandler(s *state, cmd command) error {
	feed, err := api.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return errors.New("something went wrong")
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
	}

	fmt.Println(feed)

	return nil
}
