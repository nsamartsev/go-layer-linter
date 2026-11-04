// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Nikita Samartsev

package service

import "github.com/nsamartsev/go-layer-linter/samples/testproject/internal/repository"

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{repository: userRepository}
}

func (s *UserService) RegisterUser(name string) error {
	return s.repository.Save(name)
}
