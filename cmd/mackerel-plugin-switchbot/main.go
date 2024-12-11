package main

import (
	"context"

	mp "github.com/mackerelio/go-mackerel-plugin"
	"github.com/nasa9084/go-switchbot/v4"

	"github.com/SlashNephy/mackerel-plugin-switchbot/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	plugin := mp.NewMackerelPlugin(&Plugin{
		ctx:    context.Background(),
		client: switchbot.New(cfg.SwitchBotOpenToken, cfg.SwitchBotSecretKey),
		config: cfg,
	})
	plugin.Tempfile = cfg.Tempfile
	plugin.Run()
}
