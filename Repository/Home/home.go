package Home

import "techku/Controller/Dto/response"

func (h home) GetStatus() (resp response.WebStatus, err error) {
	connection := h.dbCon.PostgreMainCon()
	query := `SELECT COUNT(1) FROM accounts.t_user_accounts`
	err = connection.QueryRow(query).Scan(&resp.RegisteredUser)
	if err != nil {
		return
	}

	query = `SELECT COUNT(1) FROM accounts.t_user_accounts_rating`
	err = connection.QueryRow(query).Scan(&resp.TotalRating)
	if err != nil {
		return
	}

	query = `SELECT COUNT(1) FROM orders.t_orders`
	err = connection.QueryRow(query).Scan(&resp.TotalOrder)
	if err != nil {
		return
	}

	query = `SELECT COUNT(1) FROM accounts.t_user_accounts tac WHERE tac.role = 'TECHNICIAN'`
	err = connection.QueryRow(query).Scan(&resp.TotalTechnician)
	if err != nil {
		return
	}

	query = `SELECT COUNT(1) FROM accounts.t_user_accounts tac WHERE tac.role = 'CUSTOMER'`
	err = connection.QueryRow(query).Scan(&resp.TotalCustomer)
	if err != nil {
		return
	}
	return
}

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
