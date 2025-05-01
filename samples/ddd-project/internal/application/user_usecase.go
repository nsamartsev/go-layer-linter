package application

import (
    "github.com/nsamartsev/ddd-project/internal/domain/repository"
)

type UserService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(id, name string) error {
    user := repository.User{ID: id, Name: name}
    return s.repo.Save(user)
}