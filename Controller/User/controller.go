package User

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"techku/Constant"
	"techku/Controller/Dto"
	"techku/Controller/Dto/request"
	"techku/Library/Helper"
)

type UserControllerInterface interface {
	Ping(g *gin.Context)
	LoginHandlerGoogle(g *gin.Context)
}

func (u user) LoginHandlerGoogle(g *gin.Context) {
	var dataRequest request.UserGoogleAuthentication
	err := g.ShouldBind(&dataRequest)
	if err != nil {
		Helper.HttpResponseError(g,
			Constant.InvalidJsonRequest.GetErrorStatus().Error,
			Constant.InvalidJsonRequest.GetErrorStatus().Code, err)
		return
	}

	claims, err := Helper.ValidateGoogleJWT(*dataRequest.GoogleJWT)
	if err != nil {
		Helper.HttpResponseError(g,
			Constant.InvalidGoogleAuth.GetErrorStatus().Error,
			Constant.InvalidGoogleAuth.GetErrorStatus().Code, err)
		return
	}
	Helper.HttpResponseSuccess(g, claims)
	//if claims.Email != user.Email {
	//	Helper.HttpResponseError(g,
	//		Constant.InvalidGoogleEmail.GetErrorStatus().Error,
	//		Constant.InvalidGoogleEmail.GetErrorStatus().Code, err)
	//	return
	//}
	return
}

func (u user) Ping(g *gin.Context) {
	fmt.Println("publishing rpc")
	Helper.HttpResponseSuccess(g, "ok")
}

func NewController(u Dto.Utilities) UserControllerInterface {
	return &user{u}
}
