# go-layer-linter

# go-layer-linter

A simple Go linter to enforce architectural layer import rules in your Go project.  
Designed for layered architectures like: `handler → service → repository`.

## 🚀 Overview

This linter ensures that certain layers of your application (e.g., `handler`, `service`, `repository`) only import allowed packages and prevent forbidden cross-layer dependencies.

## ✅ Features

- Enforces valid import paths between architectural layers
- Configurable via `.golint.yaml`
- CLI tool for easy integration into CI/CD pipelines

## 🛠️ Installation

Using Go:

```bash
go install github.com/nsamartsev/go-layer-linter/cmd/golint@latest