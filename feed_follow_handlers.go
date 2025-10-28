package main

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/mohdanas96/gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("requires url as a parameter")
	}

	feedUrl := cmd.args[0]
	_, err := url.ParseRequestURI(feedUrl)
	if err != nil {
		return fmt.Errorf("invalid url parameter")
	}

	user, err := s.db.GetUser(context.Background(), s.c.Current_user_name)
	if err != nil {
		return fmt.Errorf("something went wrong while fetching user :: %v", err)
	}

	feed, err := s.db.GetFeedWithUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("something went wrong while fetching feed :: %v", err)
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("something went wrong while creating feed follow :: %v", err)
	}

	fmt.Println(s.c.Current_user_name, "is now following", feed.Name)

	return nil
}

func handlerFollowing(s *state, _ command) error {
	user, err := s.db.GetUser(context.Background(), s.c.Current_user_name)
	if err != nil {
		return fmt.Errorf("something went wrong while fetching current user :: %v", err)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("something went wrong while retrieving feeds user follows :: %v", err)
	}

	for i, v := range feedFollows {
		fmt.Printf("%v. %v\n", i+1, v.FeedName)
	}
	return nil
}
