# syntax=docker/dockerfile:1
FROM golang:1.23.2-bookworm AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY ./ ./
RUN make build

FROM debian:bookworm-slim
WORKDIR /app

RUN <<EOF
    groupadd -g 1000 app && useradd -u 1000 -g app app
    apt-get update
    apt-get install -y --no-install-recommends ca-certificates
    apt-get clean
    rm -rf /var/lib/apt/lists/*
EOF

COPY --from=build /app/post-switchbot-metrics ./

USER app
CMD ["/app/post-switchbot-metrics"]
