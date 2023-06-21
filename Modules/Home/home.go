package Home

import "techku/Controller/Dto/response"

func (h home) GetArticles() (resp []response.ArticleList, err error) {
	resp, err = h.repo.Home.GetArticles()
	return
}

func (h home) GetStatus() (resp response.WebStatus, err error) {
	resp, err = h.repo.Home.GetStatus()
	return
}
