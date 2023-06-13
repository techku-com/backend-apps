package Order

import (
	"techku/Controller/Dto/request"
	"techku/Controller/Dto/response"
	"techku/Repository"
)

type order struct {
	repo Repository.Repository
}

type OrderModules interface {
	AddNewOrder(params request.AddNewOrder) (err error)
	GetAllOrders() (response []response.AllOrders, err error)
}

func NewModules(Repo Repository.Repository) OrderModules {
	return &order{Repo}
}
