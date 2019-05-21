package main

import (
	"fmt"
	"os"

	"github.com/axiomzen/beanstalks-api/config"
	"github.com/axiomzen/beanstalks-api/data"
	"github.com/axiomzen/beanstalks-api/server"
)

func main() {
	fmt.Println("Starting Beanstalk API...")

	// Create server
	serv := server.New(config.FromEnv())

	if os.Getenv("TEST") == "true" {
		go test(data.New(config.FromEnv()))
	}
	serv.Start()
}
