FROM golang:1.16-alpine AS builder

RUN apk --update add make

WORKDIR /services/server
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build go build -o bin/server .


FROM alpine

COPY --from=builder /services/server/bin/server /services/server/bin/server

ENTRYPOINT ["/services/server/bin/server"]
