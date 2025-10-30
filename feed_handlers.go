package main

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mohdanas96/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	args := cmd.args
	if len(args) < 2 {
		return fmt.Errorf("requires both name and url args")
	}

	feedUrl := args[len(args)-1]

	_, err := url.ParseRequestURI(feedUrl)
	if err != nil {
		return fmt.Errorf("invalid url argument")
	}

	titleName := strings.Join(args[:len(args)-1], " ")

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      titleName,
		Url:       feedUrl,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("something went wrong while creating feed :: %v", err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("something went wrong while creating feed follow when adding feed :: %v", err)
	}

	fmt.Println("Feed is successfully added")
	fmt.Println("You are now following", feed.Name)

	return nil
}

func handlerGetAllFeed(s *state, _ command) error {
	feeds, err := s.db.GetFeedsWithUsername(context.Background())
	if err != nil {
		return fmt.Errorf("something went wrong while retrieving all feeds :: %v", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found")
		return nil
	}

	for i, v := range feeds {
		fmt.Printf("%v.", i+1)
		fmt.Println("Name:", v.Name)
		fmt.Println("  URL:", v.Url)
		fmt.Println("  Username: ", v.Name_2)
		fmt.Println("  -------------------")
	}
	return nil
}
