package main

import (
	"fmt"

	"github.com/divya/resume-analyser/internal/config"
	"github.com/divya/resume-analyser/internal/logger"
	"github.com/divya/resume-analyser/internal/router"
)

func main() {

	config.Load()

	logger.Init()

	r := router.SetupRouter()

	logger.Info.Printf(
		"%s v%s started on port %s",
		config.AppConfig.AppName,
		config.AppConfig.AppVersion,
		config.AppConfig.AppPort,
	)

	err := r.Run(fmt.Sprintf(":%s", config.AppConfig.AppPort))

	if err != nil {
		logger.Error.Fatal(err)
	}
}