package main

import (
	"github.com/heCarlo/gopportunities/config"
	"github.com/heCarlo/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errf("config inicialization err: %v", err)
		return
	}

	//initialize router
	router.Initialize()
}
