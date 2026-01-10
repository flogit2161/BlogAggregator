package main

import (
	"fmt"

	"github.com/flogit2161/BlogAggregator/internal/config"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Println("Could not read from config")
	}
	err = cfg.SetUser("florian")
	if err != nil {
		fmt.Println("Could not set user")
	}
	cfg, err = config.ReadConfig()
	fmt.Println(cfg)
}
