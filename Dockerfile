FROM golang:1.21-alpine

WORKDIR /app

COPY go-layer-linter /usr/local/bin/

CMD ["go-layer-linter"]