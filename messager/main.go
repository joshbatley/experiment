package main

import (
	"github.com/rs/zerolog/log"
	utils "shared"
)

type Setting struct {
	Tps       int      `json:"tps,omitempty"`
	Tags      []string `json:"tags,omitempty"`
	LogUrl    string   `json:"logUrl,omitempty"`
	LogApiKey string   `json:"logApiKey,omitempty"`
	Env       string   `json:"env,omitempty""`
}

func main() {
	setting, err := utils.ReadConfig[Setting]()
	if err != nil {
		panic(err)
	}

	logger := utils.NewLogger().WithConsoleLogger().WithTags(setting.Tags)
	if setting.Env != "dev" {
		logger.WithHttpLogger(setting.LogApiKey, setting.LogUrl)
	}
	logger.Build()

	log.Print("Application Starting up - ", setting.LogUrl)
	defer log.Print("Application Shutting down - ", setting.LogUrl)

	NewEngine(NewInMemory(), setting.Tps).StartUp()
}
