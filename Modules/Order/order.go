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

func (o order) RateOrder(params request.RateOrder) (resp response.RateOrder, err error) {
	resp, err = o.repo.Order.RateOrder(params)
	return
}

func (o order) UpdateOrder(params request.UpdateOrder) (resp response.UpdatedOrder, err error) {
	resp, err = o.repo.Order.UpdateOrder(params)
	return
}
