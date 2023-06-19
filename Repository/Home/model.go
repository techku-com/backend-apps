package Home

import (
	"techku/Config"
	"techku/Controller/Dto/response"
)

type Repository interface {
	GetArticles() (resp []response.ArticleList, err error)
}

// CONSTRUCTOR STRUCT
type home struct {
	dbCon Config.DbConInterface
}

// CONSTRUCTOR FUNCTION FOR USER REPOSITORY
func RepositoryNew(dbCon Config.DbConInterface) Repository {
	return &home{
		dbCon: dbCon,
	}
}
