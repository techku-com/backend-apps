package User

import (
	"fmt"
	"techku/Controller/Dto/request"
	"techku/Controller/Dto/response"
)

func (u user) GetAccountInfo(params request.UserLogin) (response response.UserLogin, err error) {
	connection := u.dbCon.PostgreMainCon()
	query := `SELECT tac.password, tac.email
    		FROM accounts.t_user_accounts tac
			WHERE tac.email = $1`
	err = connection.QueryRow(query, params.Email).Scan(&response.Password, &response.Email)
	return
}

func (u user) RegisterNewAccounts(params request.UserRegistration) (response response.UserRegistration, err error) {
	connection := u.dbCon.PostgreMainCon()
	query := `INSERT INTO accounts.t_user_accounts (username, email, password)
				VALUES ($1, $2, $3) RETURNING username, email, password, TO_CHAR(created_at, 'DD-MM-YYYY HH:MI:SS')`
	err = connection.QueryRow(query, params.Username, params.Email, params.Password).Scan(&response.Username, &response.Email,
		&response.Password, &response.CreatedAt)
	return
}

func (u user) CheckExistsUser() {
	fmt.Println(u.dbCon.PostgreMainCon())
}
