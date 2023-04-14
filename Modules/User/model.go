package User

import (
	"techku/Controller/Dto/request"
	"techku/Controller/Dto/response"
	"techku/Repository"
)

type user struct {
	repo Repository.Repository
}

type UserModules interface {
	GetUser()
	RegisterNewAccounts(params request.UserRegistration) (response response.UserRegistration, err error)
	LoginAccount(params request.UserLogin) (response response.UserLogin, err error)
}

func NewModules(Repo Repository.Repository) UserModules {
	return &user{Repo}
}
