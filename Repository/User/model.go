package User

import (
	"techku/Config"
	"techku/Controller/Dto/request"
	"techku/Controller/Dto/response"
)

type Repository interface {
	CheckExistsUser()
	GetAccountInfo(params request.UserLogin) (response response.UserLogin, err error)
	RegisterNewAccounts(params request.UserRegistration) (response response.UserRegistration, err error)
}

//CONSTRUCTOR STRUCT
type user struct {
	dbCon Config.DbConInterface
}

//CONSTRUCTOR FUNCTION FOR USER REPOSITORY
func RepositoryNew(dbCon Config.DbConInterface) Repository {
	return &user{
		dbCon: dbCon,
	}
}
