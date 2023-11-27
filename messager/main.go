package main

import (
	"github.com/rs/zerolog/log"
	"messager/eventstore"
	utils "shared"
)

type Setting struct {
	Tps       int      `json:"tps,omitempty"`
	Tags      []string `json:"tags,omitempty"`
	LogUrl    string   `json:"logUrl,omitempty"`
	LogApiKey string   `json:"logApiKey,omitempty"`
	Env       string   `json:"env,omitempty"`
	AppName   string   `json:"appName,omitempty"`
}

func main() {
	setting, err := utils.ReadConfig[Setting]()
	if err != nil {
		panic(err)
	}

	logger := utils.NewLogger().WithConsoleLogger().WithTags(setting.Tags)
	if setting.Env != "dev" {
		logger.WithHttpLogger(setting.LogUrl, setting.LogApiKey)
	}
	logger.Build()

	log.Info().Msgf("Application Starting up - %s", setting.AppName)

	newRunner(eventstore.NewInMemory(), setting.Tps).startUp()
}
