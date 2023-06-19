package Modules

import (
	"techku/Modules/Home"
	"techku/Modules/Order"
	"techku/Modules/User"
)

type Modules struct {
	UserModule  User.UserModules
	OrderModule Order.OrderModules
	HomeModule  Home.HomeModules
	//repositories
}
