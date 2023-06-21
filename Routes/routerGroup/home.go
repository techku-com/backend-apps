package routerGroup

import (
	"github.com/gin-gonic/gin"
	"techku/Controller/Home"
)

type iHomeGroup interface {
	HomeApiGroup(group *gin.RouterGroup, api Home.HomeControllerInterface)
}

func (o homeGroup) HomeApiGroup(group *gin.RouterGroup, api Home.HomeControllerInterface) {
	group.GET("/articles", api.GetArticles)
	group.GET("/status", api.GetWebStatus)
}

type homeGroup struct{}

func newHomeRouterGroup() iHomeGroup {
	return &homeGroup{}
}
