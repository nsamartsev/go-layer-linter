FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /golint ./cmd/golint

FROM alpine:latest

COPY --from=builder /golint /usr/local/bin/golint

ENTRYPOINT ["golint"]
