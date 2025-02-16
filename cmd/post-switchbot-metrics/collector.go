package main

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/mackerelio/mackerel-client-go"
	"github.com/nasa9084/go-switchbot/v4"
	"github.com/samber/lo"

	"github.com/SlashNephy/mackerel-plugin-switchbot/config"
	"github.com/SlashNephy/mackerel-plugin-switchbot/metrics"
)

type Collector struct {
	config          *config.Config
	mackerelClient  *mackerel.Client
	switchBotClient *switchbot.Client
}

func (c *Collector) PostMetrics(ctx context.Context) error {
	values, err := c.fetchMetrics(ctx)
	if err != nil {
		return err
	}

	return c.mackerelClient.PostServiceMetricValues(c.config.MackerelServiceName, values)
}

func (c *Collector) fetchMetrics(ctx context.Context) ([]*mackerel.MetricValue, error) {
	devices, err := c.getDevices(ctx)
	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()
	metrics, err := metrics.FetchMetrics(ctx, c.switchBotClient, devices)
	if err != nil {
		return nil, err
	}

	var values []*mackerel.MetricValue
	for key, value := range metrics {
		deviceID, name, ok := strings.Cut(key, "-")
		if !ok {
			return nil, fmt.Errorf("invalid metric key: %s", key)
		}

		values = append(values, &mackerel.MetricValue{
			Name:  fmt.Sprintf("%s.%s.%s", c.config.Prefix, name, deviceID),
			Value: value,
			Time:  now,
		})
	}

	return values, nil
}

func (c *Collector) getDevices(ctx context.Context) ([]*switchbot.Device, error) {
	devices, _, err := c.switchBotClient.Device().List(ctx)
	if err != nil {
		return nil, err
	}

	return lo.FilterMap(devices, func(device switchbot.Device, _ int) (*switchbot.Device, bool) {
		return &device, len(c.config.DeviceIDs) == 0 || slices.Contains(c.config.DeviceIDs, device.ID)
	}), nil
}
