build: build-mackerel-plugin-switchbot build-post-switchbot-metrics

build-mackerel-plugin-switchbot:
	go build ./cmd/mackerel-plugin-switchbot

build-post-switchbot-metrics:
	go build ./cmd/post-switchbot-metrics

run: run-mackerel-plugin-switchbot

run-mackerel-plugin-switchbot:
	go run ./cmd/mackerel-plugin-switchbot

run-post-switchbot-metrics:
	go run ./cmd/post-switchbot-metrics
