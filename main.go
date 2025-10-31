package main

import (
	"database/sql"
	"log"
	"os"
	"strings"

	"github.com/mohdanas96/gator/internal/config"
	"github.com/mohdanas96/gator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	c  *config.Config
	db *database.Queries
}

func main() {
	cmdArgs := os.Args
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	configData, err := config.Read()
	if err != nil {
		log.Fatal("error while reading config")
	}

	db, err := sql.Open("postgres", configData.Db_url)
	if err != nil {
		log.Fatal("error while connecting to database")
	}

	dbQueries := database.New(db)

	configStatePtr := &state{c: &configData, db: dbQueries}

	cmds := commands{commandRegistry: make(map[string]commandHandler)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerDeleteUsers)
	cmds.register("users", handlerGetAllUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerGetAllFeed)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollowFeed))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))

	arg := cmdArgs[2:]
	cmdName := strings.TrimSpace(cmdArgs[1])

	err = cmds.run(configStatePtr, command{cmdName, arg})
	if err != nil {
		log.Fatal(err)
	}

}
