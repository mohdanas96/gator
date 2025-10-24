package main

import (
	"database/sql"
	"fmt"
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
	if len(cmdArgs) < 2 {
		fmt.Println("not enough arguments were provided")
		os.Exit(1)
	}

	configData, err := config.Read()
	if err != nil {
		fmt.Println("error while reading config", err)
	}

	db, err := sql.Open("postgres", configData.Db_url)
	if err != nil {
		fmt.Println("error while connecting to database")
	}

	dbQueries := database.New(db)

	configStatePtr := &state{c: &configData, db: dbQueries}

	cmdName := strings.TrimSpace(cmdArgs[1])

	cmds := commands{commandRegistry: make(map[string]commandHandler)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

	arg := cmdArgs[2:]
	cmd := command{cmdName, arg}

	err = cmds.run(configStatePtr, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
