package Controller

import (
	"techku/Controller/Dto"
	"techku/Controller/User"
)

type Controller struct {
	Users User.UserControllerInterface
}

func InitControllerApi(u Dto.Utilities) Controller {
	return Controller{
		Users: User.NewController(u),
	}
}
