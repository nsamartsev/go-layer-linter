BINARY := go-layer-linter

.PHONY: build test vet release clean

build:
	go build -trimpath -o $(BINARY) ./cmd/go-layer-linter

test:
	go test ./...

vet:
	go vet ./...

release:
	mkdir -p release
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -o release/$(BINARY)_linux_amd64 ./cmd/go-layer-linter
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -trimpath -o release/$(BINARY)_linux_arm64 ./cmd/go-layer-linter
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -trimpath -o release/$(BINARY)_darwin_amd64 ./cmd/go-layer-linter
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -trimpath -o release/$(BINARY)_darwin_arm64 ./cmd/go-layer-linter
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -trimpath -o release/$(BINARY)_windows_amd64.exe ./cmd/go-layer-linter

clean:
	rm -f $(BINARY) release/*
