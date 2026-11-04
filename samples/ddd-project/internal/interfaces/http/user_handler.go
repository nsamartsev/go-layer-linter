// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Nikita Samartsev

package http

import (
	"github.com/nsamartsev/go-layer-linter/samples/ddd-project/internal/application"
)

type UserHandler struct {
	useCase *application.UserUseCase
}

func NewUserHandler(useCase *application.UserUseCase) *UserHandler {
	return &UserHandler{useCase: useCase}
}
