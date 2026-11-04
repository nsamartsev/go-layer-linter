// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Nikita Samartsev

package application

import (
	"github.com/nsamartsev/go-layer-linter/samples/ddd-project/internal/domain"
	"github.com/nsamartsev/go-layer-linter/samples/ddd-project/internal/domain/repository"
)

type UserUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (u *UserUseCase) RegisterUser(id, name string) error {
	user := domain.UserEntity{ID: id, Name: name}
	return u.repo.Save(user)
}
