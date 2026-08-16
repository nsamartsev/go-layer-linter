package http

import (
    "github.com/nsamartsev/ddd-project/internal/application"
)

func StartServer() {
    service := application.NewUserService(nil)
    service.RegisterUser("1", "Alice")
}