package User

import (
	"fmt"
	"techku/Controller/Dto/request"
	"techku/Controller/Dto/response"
)

func (u user) GetAccountInfo(params request.UserLogin) (response response.UserLogin, err error) {
	connection := u.dbCon.PostgreMainCon()
	query := `SELECT tac.password, tac.email, tac.id
    		FROM accounts.t_user_accounts tac
			WHERE tac.email = $1`
	err = connection.QueryRow(query, params.Email).Scan(&response.Password, &response.Email, &response.UserId)
	return
}

func (u user) RegisterNewAccounts(params request.UserRegistration) (response response.UserRegistration, err error) {
	connection := u.dbCon.PostgreMainCon()
	query := `INSERT INTO accounts.t_user_accounts (username, email, password, role, phone_number)
				VALUES ($1, $2, $3, $4, $5) RETURNING username, email, TO_CHAR(created_at, 'DD-MM-YYYY HH:MI:SS')`
	userRole := "USER"
	if params.IsTechician {
		userRole = "TECHNICIAN"
	}
	err = connection.QueryRow(query, params.Username, params.Email, params.Password, userRole, params.Phone).Scan(&response.Username, &response.Email, &response.CreatedAt)
	return
}

func (u user) CheckExistsUser() {
	fmt.Println(u.dbCon.PostgreMainCon())
}
