# syntax=docker/dockerfile:1@sha256:2780b5c3bab67f1f76c781860de469442999ed1a0d7992a5efdf2cffc0e3d769
FROM golang:1.26.1-bookworm@sha256:ab3d6955bbc813a0f3fdf220c1d817dd89c0b3f283777db8ece4a32fe7858edd AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY ./ ./
RUN make build

FROM debian:bookworm-slim@sha256:4724b8cc51e33e398f0e2e15e18d5ec2851ff0c2280647e1310bc1642182655d
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
