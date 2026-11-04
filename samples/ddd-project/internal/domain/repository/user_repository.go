// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Nikita Samartsev

package repository

import "github.com/nsamartsev/go-layer-linter/samples/ddd-project/internal/domain"

type UserRepository interface {
	Save(user domain.UserEntity) error
}
