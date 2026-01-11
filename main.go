package main

import (
	"log"
	"os"

	"github.com/flogit2161/BlogAggregator/internal/config"
)

func main() {
	config, err := config.ReadConfig()
	if err != nil {
		log.Fatal("Could not read from config")
	}
	cfgState := &state{
		cfg: &config,
	}

	commands := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("Not enough arguments to call function")
	}

	cmdName := args[1]
	cmdArg := args[2:]

	cmd := command{
		name: cmdName,
		args: cmdArg,
	}

	err = commands.run(cfgState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
