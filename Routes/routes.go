package Routes

import (
	"github.com/gin-gonic/gin"
	"techku/Controller"
	"techku/Routes/middleware"
	"techku/Routes/routerGroup"
)

type Routes struct {
	Controller Controller.Controller
	Middleware middleware.MiddlewareInterface
	Gin        *gin.Engine
}

func (app *Routes) CollectRoutes() *gin.Engine {

	appRoute := app.Gin
	apiGroup := routerGroup.RoutesGroupCollection

	appRoute.GET("/ping", app.Controller.Users.Ping)
	//tracking api
	userGroup := appRoute.Group("/users")
	apiGroup.User.UserApiGroup(userGroup, app.Controller.Users)

	orderGroup := appRoute.Group("/orders")
	apiGroup.Order.OrderApiGroup(orderGroup, app.Controller.Orders)

	homeGroup := appRoute.Group("/home")
	apiGroup.Home.HomeApiGroup(homeGroup, app.Controller.Home)
	return appRoute
}
