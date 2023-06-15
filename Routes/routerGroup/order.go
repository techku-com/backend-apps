package routerGroup

import (
	"github.com/gin-gonic/gin"
	"techku/Controller/Order"
)

type iOrderGroup interface {
	OrderApiGroup(group *gin.RouterGroup, api Order.OrderControllerInterface)
}

func (o orderGroup) OrderApiGroup(group *gin.RouterGroup, api Order.OrderControllerInterface) {
	group.POST("/create", api.NewOrder)
	group.POST("/update", api.UpdateOrderStatus)
	group.GET("/list", api.ListOrder)
	group.POST("/rate-order", api.RateOrder)
}

type orderGroup struct{}

func newOrderRouterGroup() iOrderGroup {
	return &orderGroup{}
}
