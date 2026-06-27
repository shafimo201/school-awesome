FROM golang:1.24-alpine AS builder

WORKDIR /workspace

RUN apk add --no-cache git ca-certificates gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /workspace/bin/server ./cmd/server/main.go

FROM alpine:3.18
RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /workspace/bin/server ./server
COPY configs/config.yaml ./configs/config.yaml

EXPOSE 8080
ENTRYPOINT ["./server"]
