package main

import (
	"github.com/rs/zerolog/log"
	utils "shared"
)

type Setting struct {
	Tags    []string `json:"tags,omitempty"`
	Env     string   `json:"env,omitempty""`
	AppName string   `json:"appName,omitempty"`
}

func main() {
	setting, err := utils.ReadConfig[Setting]()
	if err != nil {
		panic(err)
	}

	logger := utils.NewLogger().WithConsoleLogger().WithTags(setting.Tags)
	logger.Build()

	log.Info().Msgf("Application Starting up - %s", setting.AppName)
	defer log.Info().Msgf("Application Shutting down - %s", setting.AppName)

	srv := NewServer()
	srv.Serve()
}
