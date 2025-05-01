package infrastructure

import (
    "github.com/nsamartsev/ddd-project/internal/interfaces/http" // ← Запрещённый импорт
)

type UserRepository struct{}

func (r *UserRepository) Save(user User) error {
    // save to DB
    return nil
}

func (r *UserRepository) FindByID(id string) (User, error) {
    // fetch from DB
    return User{}, nil
}

// User — копия из domain (дублирование!)
type User struct {
    ID   string
    Name string
}