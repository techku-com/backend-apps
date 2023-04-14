package User

import (
	"golang.org/x/crypto/bcrypt"
	"techku/Controller/Dto/request"
	"techku/Controller/Dto/response"
)

func (u user) LoginAccount(params request.UserLogin) (response response.UserLogin, err error) {
	response, err = u.repo.User.GetAccountInfo(params)
	err = bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(params.Password))
	return
}

func (u user) RegisterNewAccounts(params request.UserRegistration) (response response.UserRegistration, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	params.Password = string(hashedPassword)
	response, err = u.repo.User.RegisterNewAccounts(params)
	return
}

func (u user) GetUser() {
	u.repo.User.CheckExistsUser()
}
