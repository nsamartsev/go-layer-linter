// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Nikita Samartsev

package handler

import (
	"github.com/nsamartsev/go-layer-linter/samples/testproject/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{service: userService}
}
