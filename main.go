package main

import (
	"fmt"

	"github.com/mohdanas96/gator/internal/config"
)

func main() {
	configData, err := config.Read()
	fmt.Println("First config data ::::::::::::", configData)
	if err != nil {
		fmt.Errorf("error while reading config %v", err)
	}

	configData.Current_user_name = "Anas"
	err = configData.SetUser()
	if err != nil {
		fmt.Errorf("error while setting user to config file %v", err)
	}

	configData, err = config.Read()
	if err != nil {
		fmt.Errorf("error while reading config %v", err)
	}
	fmt.Println(configData)

}
