package Order

import (
	"techku/Controller/Dto/request"
	"techku/Controller/Dto/response"
)

func (o order) NewOrder(params request.AddNewOrder) (err error) {
	connection := o.dbCon.PostgreMainCon()
	query := `INSERT INTO orders.t_orders (created_by, issues, address, status)
				VALUES ($1, $2, $3, 1)`
	_, err = connection.Exec(query, params.UserId, params.Issues, params.Address)
	return
}

func (o order) AllOrderList() (resp []response.AllOrders, err error) {
	connection := o.dbCon.PostgreMainCon()
	query := `SELECT 
    			tac.username, 
    			tod.issues, 
    			tod.address,
    			tod.status
			FROM orders.t_orders tod
			LEFT JOIN accounts.t_user_accounts tac ON tac.id = tod.created_by
			WHERE tod.taken_by is NULL`
	rows, err := connection.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var order response.AllOrders
		err = rows.Scan(&order.CreatedBy, &order.Issues, &order.Address, &order.Status)
		if err != nil {
			return
		}
		resp = append(resp, order)
	}
	if len(resp) == 0 {
		resp = []response.AllOrders{}
	}
	return
}
