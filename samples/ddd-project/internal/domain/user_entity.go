// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Nikita Samartsev

package domain

type UserEntity struct {
	ID   string
	Name string
}

func (u UserEntity) Validate() bool {
	return u.ID != ""
}
