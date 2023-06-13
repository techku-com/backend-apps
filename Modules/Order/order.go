package Order

import (
	"techku/Controller/Dto/request"
	"techku/Controller/Dto/response"
)

func (o order) AddNewOrder(params request.AddNewOrder) (err error) {
	err = o.repo.Order.NewOrder(params)
	return
}

func (o order) GetAllOrders() (response []response.AllOrders, err error) {
	response, err = o.repo.Order.AllOrderList()
	return
}
