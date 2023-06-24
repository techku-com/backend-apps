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
    			tod.id,
    			tod.created_by,
    			tac.username, 
    			tod.issues, 
    			tod.address,
    			tod.status,
    			tac.phone_number
			FROM orders.t_orders tod
			LEFT JOIN accounts.t_user_accounts tac ON tac.id = tod.created_by
			WHERE tod.taken_by is NULL AND tod.status = 1`
	rows, err := connection.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var order response.AllOrders
		err = rows.Scan(&order.OrderId, &order.CreatedById, &order.CreatedBy, &order.Issues, &order.Address, &order.Status, &order.PhoneNumber)
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

func (o order) RateOrder(params request.RateOrder) (resp response.RateOrder, err error) {
	connection := o.dbCon.PostgreMainCon()
	query := `
		INSERT INTO accounts.t_user_accounts_rating (created_by, created_to, order_id, description, rating)
				VALUES ($1, $2, $3, $4, $5) RETURNING description, rating, order_id
	`
	err = connection.QueryRow(query, params.UserId, params.TechId, params.OrderId, params.Description, params.Rating).Scan(&resp.Description,
		&resp.Rating, &resp.OrderId)
	return
}

func (o order) UpdateOrder(params request.UpdateOrder) (resp response.UpdatedOrder, err error) {
	connection := o.dbCon.PostgreMainCon()
	query := `UPDATE orders.t_orders SET status = $1, taken_by = $4, updated_at = NOW(), price = $5, description = $6
					WHERE id = $2 AND created_by = $3 RETURNING status, id`
	err = connection.QueryRow(query, params.NewStatus, params.OrderId, params.UserId, params.TakenBy, params.Price, params.Description).Scan(&resp.Status, &resp.OrderId)
	if err != nil {
		return
	}
	if params.NewStatus == 3 {
		query = `INSERT INTO orders.t_orders_history (order_id, price, description)
				VALUES ($1, $2, $3)`
		_, err = connection.Exec(query, params.OrderId, params.Price, params.Description)
	}
	return
}
