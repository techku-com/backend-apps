package Home

import "techku/Controller/Dto/response"

func (h home) GetArticles() (resp []response.ArticleList, err error) {
	connection := h.dbCon.PostgreMainCon()
	query := `SELECT ta.title, ta.description, ta.image_url FROM articles.t_articles ta`
	rows, err := connection.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var art response.ArticleList
		err = rows.Scan(&art.Title, &art.Description, &art.Image)
		if err != nil {
			return
		}
		resp = append(resp, art)
	}
	if len(resp) == 0 {
		resp = []response.ArticleList{}
	}
	return
}
