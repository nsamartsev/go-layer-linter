// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Nikita Samartsev

package infrastructure

import "github.com/nsamartsev/go-layer-linter/samples/ddd-project/internal/domain"

type UserStore struct{}

func NewUserRepository() *UserStore {
	return &UserStore{}
}

func (r *UserStore) Save(user domain.UserEntity) error {
	// Save user in an actual persistence adapter.
	return nil
}
