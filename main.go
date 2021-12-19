package main

import (
	"Digobo/cli"
	"Digobo/config"
	"Digobo/database"
	"Digobo/log"
)

func main() {
	config.Init()
	log.Init()
	database.Init()

	err := cli.Root.Execute()
	if err != nil {
		log.Error.Fatal(err)
	}
}
