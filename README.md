# go-layer-linter

A simple Go linter to enforce architectural layer import rules in your Go project.
Designed for layered architectures like: `handler → service → repository`, or DDD projects
with `domain → application → infrastructure → interfaces` layers.

## 🚀 Overview

This linter ensures that certain layers of your application (e.g., `handler`, `service`, `repository`)
only import allowed packages and prevent forbidden cross-layer dependencies. It also checks naming
conventions (e.g. use cases must end in `UseCase`, repositories in `Repository`) and that types live
in the directory their layer expects.

## ✅ Features

- Enforces valid import paths between architectural layers
- Forbids specific packages inside the domain layer (e.g. `database/sql`, `net/http`)
- Checks naming conventions and struct location per layer
- Configurable via `.dddlint.yaml`
- CLI tool for easy integration into CI/CD pipelines

## 🛠️ Installation

Using Go:

```bash
go install github.com/nsamartsev/go-layer-linter/cmd/golint@latest
```

## 📋 Usage

Create a `.dddlint.yaml` in the directory you'll run the linter from:

```yaml
forbidden_packages_in_domain:
  - "database/sql"
  - "log"
  - "fmt"
  - "net/http"
  - "time"
  - "encoding/json"

ddd:
  layers:
    domain:
      package: "internal/domain"
      imports_allowed: []

    application:
      package: "internal/application"
      imports_allowed:
        - "internal/domain"

    infrastructure:
      package: "internal/infrastructure"
      imports_allowed:
        - "internal/domain"
        - "internal/application"

    interfaces:
      package: "internal/interfaces"
      imports_allowed:
        - "internal/domain"
        - "internal/application"
```

Then run it against a project directory:

```bash
golint run --dir path/to/your/project
```

Example, run against the DDD sample in `samples/ddd-project` (using `.golint.yaml.example` as config):

```
$ golint run --dir samples/ddd-project
[ERROR] UserService must follow the *UseCase pattern
	→ user_usecase.go
[ERROR] User must follow the *Repository pattern
	→ user_repository.go
[ERROR] Forbidden import in domain: database/sql
	→ user_entity.go
[ERROR] Repository 'UserRepository' must be in internal/domain/repository directory
	→ user_repository_impl.go

Найдено 4 нарушений.
```

Exit code is `0` when no violations are found, `1` otherwise — safe to wire into CI (see
`.github/workflow/go.yaml` for a working example).

## 🐳 Docker

```bash
docker build -t go-layer-linter .
docker run --rm -v $(pwd):/app go-layer-linter run --dir /app
```
