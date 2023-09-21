package main

import (
	"log"
	"os"
	"loanApp/pkg/config"
	"loanApp/api"
)

func main() {
	var (
		environment string
	)
	host := os.Getenv("SERVER_HOST")
	if host != "" {
		environment = "server"
	} else {
		if len(os.Args) == 2 {
			environment = os.Args[1] // developer custom file
		} else {
			environment = "local"
		}
	}
	config.Load(environment)

	if err := api.Start(); err != nil {
		log.Fatal("Failed to start server, err:", err)
		os.Exit(1)
	}
}