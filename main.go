package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/flogit2161/BlogAggregator/internal/config"
	"github.com/flogit2161/BlogAggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.ReadConfig()
	if err != nil {
		log.Fatal("Could not read from config")
	}

	db, err := sql.Open("postgres", config.DbURL)
	if err != nil {
		log.Fatal("Couldnt load database")
	}
	dbQueries := database.New(db)
	cfgState := &state{
		db:  dbQueries,
		cfg: &config,
	}

	commands := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAgg)

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
