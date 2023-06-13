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
	group.POST("/list/{status}", api.ListOrder)
}

type orderGroup struct{}

func newOrderRouterGroup() iOrderGroup {
	return &orderGroup{}
}
