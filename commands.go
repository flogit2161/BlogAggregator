package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/flogit2161/BlogAggregator/internal/config"
	"github.com/flogit2161/BlogAggregator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, exists := c.handlers[cmd.name]
	if !exists {
		return errors.New("Command doesnt exist yet")
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Println("Error fetching the next feed")
		return
	}
	_, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		fmt.Println("Error marking feed")
		return
	}
	fetchNext, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Println("Could not access feed via url")
		return
	}
	for _, f := range fetchNext.Channel.Item {
		fmt.Println(f.Title)
	}
}
