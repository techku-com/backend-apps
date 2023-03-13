package User

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"techku/Controller/Dto"
	"techku/Library/Helper"
)

type UserControllerInterface interface {
	Ping(g *gin.Context)
}

func (u user) Ping(g *gin.Context) {
	fmt.Println("publishing rpc")
	Helper.HttpResponseSuccess(g, "ok")
}

func NewController(u Dto.Utilities) UserControllerInterface {
	return &user{u}
}
