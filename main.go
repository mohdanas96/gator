package main

import (
	"fmt"

	"github.com/mohdanas96/gator/internal/config"
)

func main() {
	configData, err := config.Read()
	if err != nil {
		fmt.Println("error while reading config", err)
	}

	err = configData.SetUser("Anas")
	if err != nil {
		fmt.Println("error while setting user to config file", err)
	}

	configData, err = config.Read()
	if err != nil {
		fmt.Println("error while reading config", err)
	}

	fmt.Println(configData)
}
