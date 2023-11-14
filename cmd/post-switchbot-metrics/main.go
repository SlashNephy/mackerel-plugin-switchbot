package main

import (
	"context"
	"github.com/SlashNephy/mackerel-plugin-switchbot/config"
	"github.com/mackerelio/mackerel-client-go"
	"github.com/nasa9084/go-switchbot/v3"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	if cfg.MackerelAPIKey == "" {
		panic("MACKEREL_API_KEY is required")
	}
	if cfg.MackerelServiceName == "" {
		panic("MACKEREL_SERVICE_NAME is required")
	}

	ctx := context.Background()
	collector := &Collector{
		config:          cfg,
		mackerelClient:  mackerel.NewClient(cfg.MackerelAPIKey),
		switchBotClient: switchbot.New(cfg.SwitchBotOpenToken, cfg.SwitchBotSecretKey),
	}

	if err = collector.PostMetrics(ctx); err != nil {
		panic(err)
	}
}
