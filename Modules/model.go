package Modules

import (
	"techku/Modules/Order"
	"techku/Modules/User"
)

type Modules struct {
	UserModule  User.UserModules
	OrderModule Order.OrderModules
	//repositories
}
