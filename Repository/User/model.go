package User

import (
	"techku/Config"
)

type Repository interface {
	CheckExistsUser()
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
