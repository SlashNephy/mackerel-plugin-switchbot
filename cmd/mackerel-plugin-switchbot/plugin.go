package main

import (
	"context"
	"fmt"
	"slices"

	"github.com/SlashNephy/mackerel-plugin-switchbot/config"
	"github.com/SlashNephy/mackerel-plugin-switchbot/metrics"
	mp "github.com/mackerelio/go-mackerel-plugin"
	"github.com/nasa9084/go-switchbot/v4"
	"github.com/samber/lo"
)

type Plugin struct {
	ctx     context.Context
	client  *switchbot.Client
	config  *config.Config
	devices []*switchbot.Device
}

func (p *Plugin) MetricKeyPrefix() string {
	return p.config.Prefix
}

func (p *Plugin) GraphDefinition() map[string]mp.Graphs {
	devices, err := p.getDevices()
	if err != nil {
		panic(err)
	}

	definition := map[string]mp.Graphs{}
	for _, source := range metrics.AllMetrics {
		d := lo.Filter(devices, func(device *switchbot.Device, _ int) bool {
			sources, ok := metrics.SupportedMetrics[device.Type]
			return ok && slices.Contains(sources, source)
		})
		if len(d) == 0 {
			continue
		}

		var metrics []mp.Metrics
		for _, device := range d {
			metrics = append(metrics, mp.Metrics{
				Name:    fmt.Sprintf("%s-%s", device.ID, source.Name),
				Label:   device.Name,
				Diff:    source.Diff,
				Stacked: source.Stacked,
				Scale:   source.Scale,
			})
		}

		definition[source.Name] = mp.Graphs{
			Label:   source.Label,
			Unit:    source.Unit,
			Metrics: metrics,
		}
	}

	return definition
}

func (p *Plugin) FetchMetrics() (map[string]float64, error) {
	devices, err := p.getDevices()
	if err != nil {
		return nil, err
	}

	return metrics.FetchMetrics(p.ctx, p.client, devices)
}

func (p *Plugin) getDevices() ([]*switchbot.Device, error) {
	if p.devices == nil {
		// NOTE: infrared devices are not supported
		devices, _, err := p.client.Device().List(p.ctx)
		if err != nil {
			return nil, err
		}

		for _, device := range devices {
			if len(p.config.DeviceIDs) > 0 && !slices.Contains(p.config.DeviceIDs, device.ID) {
				continue
			}

			device := device
			p.devices = append(p.devices, &device)
		}
	}

	return p.devices, nil
}

var _ mp.PluginWithPrefix = new(Plugin)
