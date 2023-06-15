package Order

import (
	"github.com/gin-gonic/gin"
	validator "github.com/vincentius93/govalidator"
	"net/http"
	"techku/Constant"
	"techku/Controller/Dto"
	"techku/Controller/Dto/request"
	"techku/Library/Helper"
)

type OrderControllerInterface interface {
	NewOrder(g *gin.Context)
	ListOrder(g *gin.Context)
	RateOrder(g *gin.Context)
	UpdateOrderStatus(g *gin.Context)
}

func (o order) UpdateOrderStatus(g *gin.Context) {
	var dataRequest request.UpdateOrder
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
			Constant.InvalidJsonRequest.GetErrorStatus().Error,
			Constant.InvalidJsonRequest.GetErrorStatus().Code, err)
		return
	}
	resp, err := o.Modules.OrderModule.UpdateOrder(dataRequest)
	if err != nil {
		Helper.HttpResponseError(g,
			Constant.InvalidJsonRequest.GetErrorStatus().Error,
			Constant.InvalidJsonRequest.GetErrorStatus().Code, err)
		return
	}
	Helper.HttpResponseSuccess(g, resp)
	return
}

func (o order) RateOrder(g *gin.Context) {
	var dataRequest request.RateOrder
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
			Constant.InvalidJsonRequest.GetErrorStatus().Error,
			Constant.InvalidJsonRequest.GetErrorStatus().Code, err)
		return
	}
	resp, err := o.Modules.OrderModule.RateOrder(dataRequest)
	if err != nil {
		Helper.HttpResponseError(g,
			Constant.InvalidJsonRequest.GetErrorStatus().Error,
			Constant.InvalidJsonRequest.GetErrorStatus().Code, err)
		return
	}
	Helper.HttpResponseSuccess(g, resp)
	return
}

func (o order) NewOrder(g *gin.Context) {
	var dataRequest request.AddNewOrder
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
			Constant.InvalidOrder.GetErrorStatus().Error,
			Constant.InvalidOrder.GetErrorStatus().Code, err)
		return
	}
	err = o.Modules.OrderModule.AddNewOrder(dataRequest)
	if err != nil {
		Helper.HttpResponseError(g,
			Constant.InvalidOrder.GetErrorStatus().Error,
			Constant.InvalidOrder.GetErrorStatus().Code, err)
		return
	}
	Helper.HttpResponseSuccess(g, http.StatusOK)
	return
}

func (o order) ListOrder(g *gin.Context) {
	resp, err := o.Modules.OrderModule.GetAllOrders()
	if err != nil {
		Helper.HttpResponseError(g,
			Constant.InvalidOrder.GetErrorStatus().Error,
			Constant.InvalidOrder.GetErrorStatus().Code, err)
		return
	}
	Helper.HttpResponseSuccess(g, resp)
	return
}

func NewController(u Dto.Utilities) OrderControllerInterface {
	return &order{u}
}
