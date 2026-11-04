# Repository Guidelines

## Project Structure & Module Organization

This repository ships a standalone Go CLI. The executable entry point is `cmd/go-layer-linter`; command handling, configuration parsing, and rule analysis live in `internal/cli`, `internal/config`, and `internal/rules`. These `internal/` packages are implementation details, not a supported library API. Unit tests sit beside their packages in `*_test.go` files. `samples/` contains valid and intentionally invalid projects for exercising the CLI. The root `.dddlint.yaml.example` documents configuration; update it when config behavior changes.

## Build, Test, and Development Commands

- `make build` compiles `./cmd/go-layer-linter` into `go-layer-linter`.
- `make test` runs `go test ./...`.
- `make vet` runs `go vet ./...`.
- `go run ./cmd/go-layer-linter run --dir samples/ddd-project` runs the DDD sample; use `samples/testproject` for the handler → service → repository example.
- `make release` builds the configured Linux, macOS, and Windows binaries under `release/`.

## Coding Style & Naming

Format Go changes with `gofmt`; use lowercase package names and descriptive exported identifiers. Keep CLI wiring in `internal/cli`, config validation in `internal/config`, and analysis rules in `internal/rules`. Go source files use the MIT SPDX and copyright header pattern already present in the repository.

## Testing Guidelines

Use Go's standard `testing` package. Name tests `Test…` in files ending `_test.go`; prefer table-driven cases for rule and config behavior. Add or update tests when behavior changes, and run `make test` and `make vet` before opening a pull request. There is no published coverage threshold.

## Commits & Pull Requests

Recent history uses short, imperative commit subjects (for example, “Prepare project for open-source release”); follow that concise style. A pull request should describe the behavior changed and include output from relevant checks. Update the README or configuration example for user-visible changes, and follow `CONTRIBUTING.md` and `CODE_OF_CONDUCT.md`.

## Security & Configuration

Use `.dddlint.yaml.example` and the sample project configs as generic configuration examples; keep organization-specific names, URLs, and credentials out of defaults and fixtures. Report vulnerabilities privately as described in `SECURITY.md`.
