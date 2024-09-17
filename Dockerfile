# syntax=docker/dockerfile:1
FROM golang:1.23.1-bookworm AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN make build

FROM debian:bookworm@sha256:133a1f2aa9e55d1c93d0ae1aaa7b94fb141265d0ee3ea677175cdb96f5f990e5
WORKDIR /app

RUN groupadd -g 1000 app && useradd -u 1000 -g app app \
    && apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=build /app/post-switchbot-metrics ./

USER app
CMD ["./post-switchbot-metrics"]
