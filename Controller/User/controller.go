package User

import (
	"fmt"
	"github.com/gin-gonic/gin"
	validator "github.com/vincentius93/govalidator"
	"strconv"
	"techku/Constant"
	"techku/Controller/Dto"
	"techku/Controller/Dto/request"
	"techku/Library/Helper"
)

type UserControllerInterface interface {
	Ping(g *gin.Context)
	LoginHandlerGoogle(g *gin.Context)
	RegisterAccount(g *gin.Context)
	LoginAccount(g *gin.Context)
	MyOrderList(g *gin.Context)
}

func (u user) MyOrderList(g *gin.Context) {
	userId, err := strconv.Atoi(g.Param("user_id"))
	if err != nil {
		Helper.HttpResponseError(g,
			Constant.InvalidJsonRequest.GetErrorStatus().Error,
			Constant.InvalidJsonRequest.GetErrorStatus().Code, err)
		return
	}

	response, err := u.Modules.UserModule.MyOrderList(userId)
	if err != nil {
		Helper.HttpResponseError(g,
			Constant.InvalidJsonRequest.GetErrorStatus().Error,
			Constant.InvalidJsonRequest.GetErrorStatus().Code, err)
		return
	}
	Helper.HttpResponseSuccess(g, response)
}

func (u user) LoginAccount(g *gin.Context) {
	var dataRequest request.UserLogin
	err := g.ShouldBind(&dataRequest)
	if err != nil {
		Helper.HttpResponseError(g,
			Constant.InvalidJsonRequest.GetErrorStatus().Error,
			Constant.InvalidJsonRequest.GetErrorStatus().Code, err)
		return
	}

	response, err := u.Modules.UserModule.LoginAccount(dataRequest)
	if err != nil {
		Helper.HttpResponseError(g,
			Constant.InvalidLogin.GetErrorStatus().Error,
			Constant.InvalidLogin.GetErrorStatus().Code, err)
		return
	}
	Helper.HttpResponseSuccess(g, response)
	return
}

func (u user) RegisterAccount(g *gin.Context) {
	var dataRequest request.UserRegistration
	err := g.ShouldBind(&dataRequest)
	if err != nil {
		Helper.HttpResponseError(g,
			Constant.InvalidJsonRequest.GetErrorStatus().Error,
			Constant.InvalidJsonRequest.GetErrorStatus().Code, err)
		return
	}
	err = validator.Validate(dataRequest)
	if err != nil {
		Helper.HttpResponseError(g,
			Constant.InvalidRegister.GetErrorStatus().Error,
			Constant.InvalidRegister.GetErrorStatus().Code, err)
		return
	}
	response, err := u.Modules.UserModule.RegisterNewAccounts(dataRequest)
	if err != nil {
		Helper.HttpResponseError(g,
			Constant.InvalidRegister.GetErrorStatus().Error,
			Constant.InvalidRegister.GetErrorStatus().Code, err)
		return
	}
	Helper.HttpResponseSuccess(g, response)
	return
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
