FROM golang:1.16-alpine AS builder

RUN apk --update add make

WORKDIR /services/db-manager
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build make build


FROM alpine

RUN GRPC_HEALTH_PROBE_VERSION=v0.3.2 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

COPY --from=builder /services/db-manager/bin/server /services/db-manager/bin/server

ENTRYPOINT ["/services/db-manager/bin/server"]
