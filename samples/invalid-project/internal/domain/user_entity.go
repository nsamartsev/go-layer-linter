// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Nikita Samartsev

package domain

import "net/http"

type UserEntity struct {
	client *http.Client
}
