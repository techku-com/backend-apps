package Controller

import (
	"techku/Controller/Dto"
	"techku/Controller/Order"
	"techku/Controller/User"
)

type Controller struct {
	Users  User.UserControllerInterface
	Orders Order.OrderControllerInterface
}

func InitControllerApi(u Dto.Utilities) Controller {
	return Controller{
		Users:  User.NewController(u),
		Orders: Order.NewController(u),
	}
}
