package metrics

import (
	"context"
	"fmt"
	"sync"

	"github.com/nasa9084/go-switchbot/v4"
	"golang.org/x/sync/errgroup"
)

func FetchMetrics(ctx context.Context, client *switchbot.Client, devices []*switchbot.Device) (map[string]float64, error) {
	results := map[string]float64{}

	var mutex sync.Mutex
	eg, egctx := errgroup.WithContext(ctx)
	for _, device := range devices {
		sources, ok := SupportedMetrics[device.Type]
		if !ok {
			continue
		}

		device := device
		eg.Go(func() error {
			status, err := client.Device().Status(egctx, device.ID)
			if err != nil {
				return err
			}

			mutex.Lock()
			defer mutex.Unlock()

			for _, source := range sources {
				key := fmt.Sprintf("%s-%s", device.ID, source.Name)
				results[key] = source.Value(&status)
			}

			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return results, nil
}
