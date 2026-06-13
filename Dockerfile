# syntax=docker/dockerfile:1@sha256:87999aa3d42bdc6bea60565083ee17e86d1f3339802f543c0d03998580f9cb89
FROM golang:1.26.4-bookworm@sha256:5f68ec6805843bd3981a951ffada82a26a0bd2631045c8f7dba483fa868f5ec5 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY ./ ./
RUN make build

FROM debian:bookworm-slim@sha256:96e378d7e6531ac9a15ad505478fcc2e69f371b10f5cdf87857c4b8188404716
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
