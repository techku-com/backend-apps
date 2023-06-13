package Repository

import (
	"techku/Config"
	"techku/Repository/Order"
	"techku/Repository/User"
)

// CONSTRUCTOR STRUCT FOR ALL REPOSITORY
type Repository struct {
	User  User.Repository
	Order Order.Repository
}

// REPOSITORY INITIALIZATION
func InitRepo(dbCon Config.DbConInterface) Repository {
	return Repository{
		User:  User.RepositoryNew(dbCon),
		Order: Order.RepositoryNew(dbCon),
	}
}
