// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Nikita Samartsev

package repository

type UserRepository interface {
	Save(name string) error
}
