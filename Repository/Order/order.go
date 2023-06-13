package Order

import (
	"techku/Controller/Dto/request"
	"techku/Controller/Dto/response"
)

func (o order) NewOrder(params request.AddNewOrder) (err error) {
	connection := o.dbCon.PostgreMainCon()
	query := `INSERT INTO orders.t_orders (created_by, description, location, contact_person, hardware_type, service_type, scheduled_date)
				VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = connection.Exec(query, params.UserId, params.Description, params.Location,
		params.ContactPerson, params.HardwareType, params.ServiceType, params.ScheduleDate)
	return
}

func (o order) AllOrderList() (resp []response.AllOrders, err error) {
	connection := o.dbCon.PostgreMainCon()
	query := `SELECT 
    			created_by, 
    			description, 
    			location, 
    			contact_person, 
    			hardware_type, 
    			service_type, 
    			schedule_date
			FROM orders.t_orders to WHERE to.taken_by = null`
	rows, err := connection.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var order response.AllOrders
		err = rows.Scan(&order.CreatedBy, &order.Description, &order.Location, &order.ContactPerson, &order.HardwareType, &order.ServiceType, &order.ScheduleDate)
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
