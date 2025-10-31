package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/mohdanas96/gator/internal/database"
)

var queryLimit int32

func handlerBrowse(s *state, cmd command, user database.User) error {
	queryLimit = 2
	if len(cmd.args) > 0 {
		num, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("invalid limit :: %v", err)
		}
		queryLimit = int32(num)
	}

	posts, err := s.db.GetPostsUser(context.Background(), database.GetPostsUserParams{UserID: user.ID, Limit: queryLimit})
	if err != nil {
		return fmt.Errorf("cannot fetch posts :: %v", err)
	}

	if len(posts) == 0 {
		fmt.Println("You are not following any feeds")
	}

	for i, v := range posts {
		fmt.Println("----------------------------------------------")
		fmt.Printf("%v. Post Title: %v\n", i+1, v.Title)
		fmt.Printf("    Feed: %v: %v\n", v.FeedName, v.FeedUrl)
		// fmt.Printf("    Description: %v\n", v.Description.String)
		fmt.Printf("    Published on: %v.%v.%v\n", v.PublishedAt.Day(), v.PublishedAt.Month(), v.PublishedAt.Year())
		fmt.Printf("    Url to full content: %v\n", v.Url)
	}

	return nil
}
