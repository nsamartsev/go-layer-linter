// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Nikita Samartsev

package main

import (
	"github.com/nsamartsev/go-layer-linter/samples/ddd-project/internal/application"
	"github.com/nsamartsev/go-layer-linter/samples/ddd-project/internal/infrastructure"
	"github.com/nsamartsev/go-layer-linter/samples/ddd-project/internal/interfaces/http"
)

func main() {
	repository := infrastructure.NewUserRepository()
	useCase := application.NewUserUseCase(repository)
	_ = http.NewUserHandler(useCase)
}
