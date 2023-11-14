package config

import (
	"github.com/jessevdk/go-flags"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	SwitchBotOpenToken  string   `env:"SWITCHBOT_OPEN_TOKEN" long:"open-token" short:"t" description:"SwitchBot Open Token" required:"true"`
	SwitchBotSecretKey  string   `env:"SWITCHBOT_SECRET_KEY" long:"secret-key" short:"k" description:"SwitchBot Secret Key" required:"true"`
	DeviceIDs           []string `env:"SWITCHBOT_DEVICE_IDS" env-delim:"," long:"device-id" short:"d" description:"Device ID(s) to collect metrics"`
	Prefix              string   `env:"PREFIX" long:"prefix" description:"Metric Key Prefix" default:"switchbot"`
	Tempfile            string   `env:"TEMPFILE" long:"tempfile" description:"Temp Filename"`
	MackerelAPIKey      string   `env:"MACKEREL_API_KEY" long:"mackerel-api-key" description:"Mackerel API Key"`
	MackerelServiceName string   `env:"MACKEREL_SERVICE_NAME" long:"mackerel-service-name" description:"Mackerel Service Name"`
}

func LoadConfig() (*Config, error) {
	var config Config
	parser := flags.NewParser(&config, flags.Default)
	if _, err := parser.Parse(); err != nil {
		return nil, err
	}

	return &config, nil
}
