package User

import (
	"techku/Repository"
)

type user struct {
	repo Repository.Repository
}

type UserModules interface {
	GetUser()
}

func NewModules(Repo Repository.Repository) UserModules {
	return &user{Repo}
}
