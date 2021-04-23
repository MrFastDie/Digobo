package main

import (
	"Digobo/cli"
	"Digobo/config"
	"Digobo/log"
)

func main() {
	config.Init()
	log.Init()

	err := cli.Root.Execute()
	if err != nil {
		log.Error.Fatal(err)
	}
}