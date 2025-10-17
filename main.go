package main

import (
	"fmt"
	"os"

	"github.com/mohdanas96/gator/internal/config"
)

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

	configStatePtr := &state{c: &configData}

	cmdName := cmdArgs[0]

	cmds := commands{commandRegistry: make(map[string]commandHandler)}
	cmds.register(cmdName, handlerLogin)

	arg := cmdArgs[2:]
	cmd := command{cmdName, arg}

	err = cmds.run(configStatePtr, cmd)
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}

}
