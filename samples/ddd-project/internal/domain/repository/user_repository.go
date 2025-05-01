package repository

type UserRepository interface {
	Save(user User) error
	FindByID(id string) (User, error)
}

// Нарушение: не должно быть здесь
type User struct {
	ID   string
	Name string
}
