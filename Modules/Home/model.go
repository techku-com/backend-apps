package Home

import (
	"techku/Controller/Dto/response"
	"techku/Repository"
)

type home struct {
	repo Repository.Repository
}

type HomeModules interface {
	GetArticles() (resp []response.ArticleList, err error)
	GetStatus() (resp response.WebStatus, err error)
}

func NewModules(Repo Repository.Repository) HomeModules {
	return &home{Repo}
}
