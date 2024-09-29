package main

import (
	"github.com/marcospsw/first-project-go/config"
	"github.com/marcospsw/first-project-go/router"
)

func main() {
	// Catch Logger
	logger := config.GetLogger()

	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errf("Initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
