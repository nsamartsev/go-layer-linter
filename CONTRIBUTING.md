# Contributing

Thanks for considering a contribution.

This project is currently maintained as a standalone CLI. Its `internal/` packages are not a supported library interface.

## Before opening a pull request

1. Keep changes focused and include tests for behavior changes.
2. Format Go code with `gofmt`.
3. Run `go test ./...`, `go vet ./...`, and both sample commands from the README.
4. Update the README or configuration example when user-visible behavior changes.

Pull requests should explain the behavior being changed and include the command output for relevant checks.

## Local checks

```sh
make test
make vet
go run ./cmd/go-layer-linter run --dir samples/ddd-project
go run ./cmd/go-layer-linter run --dir samples/testproject
```

Releases use SemVer tags in `vMAJOR.MINOR.PATCH` form, matching the `v*` release workflow trigger. Use `v0.x` while the CLI behavior and configuration are still settling; the planned first release is `v0.1.0`. The workflow builds archives for Linux, macOS, and Windows and publishes SHA-256 checksums.

Please follow the [Code of Conduct](CODE_OF_CONDUCT.md). Report security issues privately as described in [SECURITY.md](SECURITY.md).
