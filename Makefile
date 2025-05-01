BINARY=ddd-linter
VERSION=$(shell git describe --tags 2>/dev/null || echo "dev")

build:
    go build -o ${BINARY} cmd/golint/main.go

release:
    GOOS=linux   GOARCH=amd64 go build -o release/${BINARY}_linux_amd64
    GOOS=darwin  GOARCH=amd64 go build -o release/${BINARY}_darwin_amd64
    GOOS=windows GOARCH=amd64 go build -o release/${BINARY}_windows_amd64.exe

clean:
    rm -f ${BINARY} release/*