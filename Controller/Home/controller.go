package Home

import (
	"github.com/gin-gonic/gin"
	"techku/Constant"
	"techku/Controller/Dto"
	"techku/Library/Helper"
)

type HomeControllerInterface interface {
	GetArticles(g *gin.Context)
	GetWebStatus(g *gin.Context)
}

func (h home) GetWebStatus(g *gin.Context) {
	response, err := h.Modules.HomeModule.GetStatus()
	if err != nil {
		Helper.HttpResponseError(g,
			Constant.InvalidJsonRequest.GetErrorStatus().Error,
			Constant.InvalidJsonRequest.GetErrorStatus().Code, err)
		return
	}
	Helper.HttpResponseSuccess(g, response)
	return
}

func (h home) GetArticles(g *gin.Context) {
	response, err := h.Modules.HomeModule.GetArticles()
	if err != nil {
		Helper.HttpResponseError(g,
			Constant.InvalidJsonRequest.GetErrorStatus().Error,
			Constant.InvalidJsonRequest.GetErrorStatus().Code, err)
		return
	}
	Helper.HttpResponseSuccess(g, response)
	return
}

func NewController(u Dto.Utilities) HomeControllerInterface {
	return &home{u}
}
