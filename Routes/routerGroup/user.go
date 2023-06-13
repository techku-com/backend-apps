package routerGroup

import (
	"github.com/gin-gonic/gin"
	"techku/Controller/User"
)

type iUserGroup interface {
	UserApiGroup(group *gin.RouterGroup, api User.UserControllerInterface)
}

func (r userGroup) UserApiGroup(group *gin.RouterGroup, api User.UserControllerInterface) {
	group.POST("/authenticate", api.LoginHandlerGoogle)
	group.POST("/register", api.RegisterAccount)
	group.POST("/login", api.LoginAccount)
}

type userGroup struct{}

func newUserRouterGroup() iUserGroup {
	return &userGroup{}
}
