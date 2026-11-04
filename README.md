# go-layer-linter

`go-layer-linter` is a standalone CLI for checking package imports in DDD-style layered Go projects. Configure package paths and allowed imports in `.dddlint.yaml`; optional naming and type-location conventions can be disabled there as well.

The project is currently CLI-only. The packages under `internal/` are implementation details and are not supported as an importable library API.

## Requirements

- Go 1.26 or newer
- A `go.mod` file at the root of the project being analyzed
- A `.dddlint.yaml` file in the root of the project being analyzed

## Install

```sh
go install github.com/nsamartsev/go-layer-linter/cmd/go-layer-linter@latest
```

The release workflow publishes Linux, macOS, and Windows archives with SHA-256 checksums when a `v*` tag is pushed. See [Releases](https://github.com/nsamartsev/go-layer-linter/releases).

## Configure

Copy [`.dddlint.yaml.example`](.dddlint.yaml.example) to `.dddlint.yaml` in the project you want to analyze, then edit the package paths and allowed imports for its structure.

```yaml
forbidden_packages_in_domain:
  - database/sql
  - net/http

ddd:
  layers:
    domain:
      package: internal/domain
      imports_allowed: []
    application:
      package: internal/application
      imports_allowed:
        - internal/domain
    infrastructure:
      package: internal/infrastructure
      imports_allowed:
        - internal/domain
        - internal/application
    interfaces:
      package: internal/interfaces
      imports_allowed:
        - internal/application
```

Layer package paths are relative to the analyzed project root. The analyzer combines each path with the
module path in `go.mod`, so `internal/domain` matches an import such as
`github.com/example/project/internal/domain`. Entries in `imports_allowed` may be relative layer
paths or full import paths. A layer may import its own subpackages.

Naming and type-location checks are enabled when omitted from `rules`. Set individual entries to `false` to opt out. Repository placement applies to repository interfaces; concrete adapters can keep names such as `UserRepository` in infrastructure packages.

For example, disable the service naming convention with:

```yaml
rules:
  naming:
    service: false
```

## Run

```sh
go-layer-linter run --dir path/to/project
```

The linter reads `path/to/project/.dddlint.yaml` and the module path from
`path/to/project/go.mod` to identify project imports. It exits with status `0` when no violations are
found and `1` for violations or analysis errors. It skips `.git`, `vendor`, `testdata`, `examples`,
`fixtures`, and `samples` directories.

The repository includes two runnable examples:

```sh
go-layer-linter run --dir samples/ddd-project
go-layer-linter run --dir samples/testproject
```

The first demonstrates DDD package rules. The second demonstrates a handler → service → repository import chain.
The intentionally invalid `samples/invalid-project` fixture demonstrates violation output and exit status `1`.

## Checks

- Imports between configured layers, including full module import paths
- Forbidden imports in the configured `domain` layer
- Optional `*Repository`, `*Service`, and `*UseCase` suffixes in `*_repository.go`, `*_service.go`, and `*_usecase.go` files
- Optional `*Entity` placement under the configured domain package
- Optional repository-interface placement under `<domain package>/repository`
- Optional `*UseCase` placement in `*_usecase.go` files

## Development

```sh
make test
make vet
make build
```

See [CONTRIBUTING.md](CONTRIBUTING.md) for the local workflow and [THIRD_PARTY_NOTICES.md](THIRD_PARTY_NOTICES.md) for dependency attribution.

Community expectations are in [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md); vulnerability reporting is in [SECURITY.md](SECURITY.md).

## Docker

```sh
docker build -t go-layer-linter .
docker run --rm -v "$(pwd)/samples/ddd-project:/app" go-layer-linter run --dir /app
```

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE).
