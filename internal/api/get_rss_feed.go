package api

import (
	"context"
	"encoding/xml"
	"errors"
	"html"
	"io"
	"net/http"
)

func FetchFeed(context context.Context, feedUrl string) (*RSSFeed, error) {
	if len(feedUrl) < 1 {
		return nil, errors.New("rssFeed url must be specified")
	}

	req, err := http.NewRequestWithContext(context, "GET", feedUrl, nil)
	if err != nil {
		return nil, errors.New("something happened while creating request")
	}

	req.Header.Set("User-Agent", "gator")

	c := http.Client{}

	res, err := c.Do(req)
	if err != nil {
		return nil, errors.New("something went wrong while fetching feed")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("something went wrong while reading response body")
	}

	rssFeed := RSSFeed{}
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return nil, errors.New("something went wrong while parsing xml")
	}

	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)

	for i := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[i].Description = html.UnescapeString(rssFeed.Channel.Item[i].Description)
		rssFeed.Channel.Item[i].Title = html.UnescapeString(rssFeed.Channel.Item[i].Title)
	}

	return &rssFeed, nil
}
