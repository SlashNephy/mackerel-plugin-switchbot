package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/jessevdk/go-flags"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	SwitchBotOpenToken string   `long:"open-token" description:"SwitchBot Open Token" required:"true" env:"SWITCHBOT_OPEN_TOKEN"`
	SwitchBotSecretKey string   `long:"secret-key" description:"SwitchBot Secret Key" required:"true" env:"SWITCHBOT_SECRET_KEY"`
	FilterDevices      []string `long:"filter-devices" description:"Filter devices by ID" env:"FILTER_DEVICES"`
	Prefix             string   `long:"prefix" description:"Metric key prefix" default:"switchbot" env:"PREFIX"`
	Tempfile           string   `long:"tempfile" description:"Temp filename" env:"TEMPFILE"`
}

func LoadConfig() (*Config, error) {
	var config Config
	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	parser := flags.NewParser(&config, flags.Default)
	if _, err := parser.Parse(); err != nil {
		return nil, err
	}

	return &config, nil
}
