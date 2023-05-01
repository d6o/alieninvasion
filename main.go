package main

import (
	"context"
	"log"
	"os"

	"github.com/d6o/alieninvasion/internal/app"
	"github.com/d6o/alieninvasion/internal/cli"
)

const (
	statusErrorCode = 1
)

func main() {
	log.Println("Starting Alien Invasion")

	ctx := context.Background()

	cfg, err := cli.Parse()
	if err != nil {
		log.Println("Error parsing cli flags:", err)
		os.Exit(statusErrorCode)
	}

	application := app.NewApp(cfg.Aliens, cfg.Input, cfg.Output, cfg.Verbose)

	if err := application.Run(ctx); err != nil {
		log.Println("Error while executing application:", err)
		os.Exit(statusErrorCode)
	}

	log.Println("Invasion completed")
}
