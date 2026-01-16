package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/flogit2161/BlogAggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("Please only use 1 argument to use command function")
	}

	name := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		log.Fatal("User doesnt exist")
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return errors.New("Error setting up the user's name")
	}

	fmt.Println("User has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("Please use at least 1 argument to register")
	}
	name := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		log.Fatal("User already exists")
	}

	newClient, err := s.db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      name,
		},
	)
	if err != nil {
		return errors.New("Error creating the new user")
	}
	err = s.cfg.SetUser(newClient.Name)
	if err != nil {
		return errors.New("Error setting user's name in config")
	}
	fmt.Println("User created sucessfully")
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return errors.New("Error reseting the users table")
	}
	fmt.Println("Users table successfully reseted")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return errors.New("Could not get users from table")
	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %s\n", user.Name)
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {
	timeString := cmd.args[0]
	time_between_reqs, err := time.ParseDuration(timeString)
	if err != nil {
		return errors.New("Could not convert arg to time duration")
	}
	fmt.Printf("Collecting feeds every %v\n Please press CTRL + C to kill program", time_between_reqs)
	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return errors.New("Please enter at least 2 arguments for command to work (name, url)")
	}
	name := cmd.args[0]
	url := cmd.args[1]

	feed, err := s.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      name,
			Url:       url,
			UserID:    user.ID,
		},
	)
	if err != nil {
		return errors.New("Error creating feed")
	}
	fmt.Println("Successfully created feed")

	follow, err := s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: feed.CreatedAt,
			UpdatedAt: feed.UpdatedAt,
			UserID:    feed.UserID,
			FeedID:    feed.ID,
		},
	)
	if err != nil {
		return errors.New("Could not assign feed to users following")
	}
	fmt.Println("Feed added to user's follows")
	fmt.Printf("User ID : %s\n", follow.UserID)
	fmt.Printf("Feed name : %s\n", follow.FeedName)
	return nil

}

func handlerFeed(s *state, cmd command) error {
	feeds, err := s.db.GetFeedsWithUsers(context.Background())
	if err != nil {
		return errors.New("Could not access feeds in db")
	}
	for _, f := range feeds {
		fmt.Printf("%s\n%s\n%s\n", f.Name, f.Url, f.Name_2)
	}
	return nil
}

func handlerFollow(s *state, cmd command, user database.User) error {
	url := cmd.args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return errors.New("Couldnt access feed via URL")
	}

	followedFeed, err := s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    user.ID,
			FeedID:    feed.ID,
		},
	)
	if err != nil {
		return errors.New("Could not create feed")
	}
	fmt.Println("Feed record created successfully")
	fmt.Printf("Username : %s\n", followedFeed.UserName)
	fmt.Printf("Feedname : %s\n", followedFeed.FeedName)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return errors.New("Could not access followed feeds")
	}
	for _, f := range follows {
		fmt.Println(f.FeedName)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	url := cmd.args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return errors.New("Could not access feed via url")
	}

	err = s.db.UnfollowFeed(
		context.Background(),
		database.UnfollowFeedParams{
			FeedID: feed.ID,
			UserID: user.ID,
		},
	)
	fmt.Println("Successfully unfollowed feed")
	return nil
}
