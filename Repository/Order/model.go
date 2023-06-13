package Order

import (
	"techku/Config"
	"techku/Controller/Dto/request"
	"techku/Controller/Dto/response"
)

type Repository interface {
	NewOrder(params request.AddNewOrder) (err error)
	AllOrderList() (resp []response.AllOrders, err error)
}

// CONSTRUCTOR STRUCT
type order struct {
	dbCon Config.DbConInterface
}

// CONSTRUCTOR FUNCTION FOR USER REPOSITORY
func RepositoryNew(dbCon Config.DbConInterface) Repository {
	return &order{
		dbCon: dbCon,
	}
}
