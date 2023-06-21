package User

import (
	"fmt"
	"techku/Controller/Dto/request"
	"techku/Controller/Dto/response"
)

func (u user) GetAccountInfo(params request.UserLogin) (response response.UserLogin, err error) {
	connection := u.dbCon.PostgreMainCon()
	query := `SELECT tac.password, tac.email, tac.id, tac.role, tac.username
    		FROM accounts.t_user_accounts tac
			WHERE tac.email = $1`
	err = connection.QueryRow(query, params.Email).Scan(&response.Password, &response.Email, &response.UserId, &response.Role, &response.Username)
	return
}

func (u user) RegisterNewAccounts(params request.UserRegistration) (response response.UserRegistration, err error) {
	connection := u.dbCon.PostgreMainCon()
	query := `INSERT INTO accounts.t_user_accounts (username, email, password, role, phone_number)
				VALUES ($1, $2, $3, $4, $5) RETURNING username, email, TO_CHAR(created_at, 'DD-MM-YYYY HH:MI:SS')`
	userRole := "CUSTOMER"
	if params.Role {
		userRole = "TECHNICIAN"
	}
	err = connection.QueryRow(query, params.Username, params.Email, params.Password, userRole, params.Phone).Scan(&response.Username, &response.Email, &response.CreatedAt)
	return
}

func (u user) MyOrderList(userId int) (resp []response.MyOrderList, err error) {
	connection := u.dbCon.PostgreMainCon()
	query := `SELECT 
    			tac.username, 
    			COALESCE(tac2.username, ''),
    			tac.id,
    			COALESCE(tac2.id, 0),
    			TO_CHAR(tod.created_at, 'DD-MM-YYYY'),
    			tod.issues, 
    			tod.status,
    			COALESCE(tacr.rating, 0),
    			COALESCE(tacr.description, '')
			FROM orders.t_orders tod
			LEFT JOIN accounts.t_user_accounts tac ON tac.id = tod.created_by
			LEFT JOIN accounts.t_user_accounts tac2 on tac2.id = tod.taken_by
			LEFT JOIN accounts.t_user_accounts_rating tacr on tacr.order_id = tod.id
			WHERE tod.created_by = $1 OR tod.taken_by = $1`
	rows, err := connection.Query(query, userId)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var order response.MyOrderList
		err = rows.Scan(&order.CreatedByName, &order.TakenByName, &order.CreatedBy, &order.TakenBy, &order.CreatedAt, &order.Issues, &order.Status,
			&order.Rating.Rating, &order.Rating.Description)
		if err != nil {
			return
		}
		resp = append(resp, order)
	}
	if len(resp) == 0 {
		resp = []response.MyOrderList{}
	}
	return
}

func (u user) CheckExistsUser() {
	fmt.Println(u.dbCon.PostgreMainCon())
}
