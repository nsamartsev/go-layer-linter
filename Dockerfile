FROM golang:1.27-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -trimpath -o /go-layer-linter ./cmd/go-layer-linter

FROM alpine:latest

COPY --from=builder /go-layer-linter /usr/local/bin/go-layer-linter
COPY LICENSE /usr/share/licenses/go-layer-linter/LICENSE
COPY THIRD_PARTY_NOTICES.md /usr/share/doc/go-layer-linter/THIRD_PARTY_NOTICES.md
COPY licenses/ /usr/share/licenses/go-layer-linter/third-party/

ENTRYPOINT ["go-layer-linter"]
