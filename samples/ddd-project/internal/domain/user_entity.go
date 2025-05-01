package domain

import (
    "database/sql"
)

type User struct {
    ID   string
    Name string
}

func (u *User) Validate() bool {
    return u.ID != ""
}