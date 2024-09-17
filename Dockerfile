# syntax=docker/dockerfile:1
FROM golang:1.23.1-bookworm AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN make build

FROM debian:bookworm@sha256:b8084b1a576c5504a031936e1132574f4ce1d6cc7130bbcc25a28f074539ae6b
WORKDIR /app

RUN groupadd -g 1000 app && useradd -u 1000 -g app app \
    && apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=build /app/post-switchbot-metrics ./

USER app
CMD ["./post-switchbot-metrics"]
