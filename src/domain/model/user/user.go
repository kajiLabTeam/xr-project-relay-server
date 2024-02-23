package user_model_domain

import (
	"fmt"

	"github.com/kajiLabTeam/xr-project-relay-server/utils"
)

type User struct {
	id string
}

func NewUser(id string) (*User, error) {
	if !utils.IsValidURL(id) {
		return nil, fmt.Errorf("invalid id value")
	}
	return &User{id: id}, nil
}

func (u *User) GetId() string {
	return u.id
}
