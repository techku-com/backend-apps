package Helper

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"techku/Constant"
	"time"
)

type ResponsePublisher struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Code   int         `json:"code"`
}

type setResponse struct {
	Status     string      `json:"status"`
	Data       interface{} `json:"data"`
	AccessTime string      `json:"accessTime"`
}

func HttpResponseSuccess(g *gin.Context, data interface{}) {
	location, _ := time.LoadLocation(Constant.TimeLocation)
	responseData := setResponse{
		Status:     http.StatusText(http.StatusOK),
		Data:       data,
		AccessTime: time.Now().In(location).Format("02-01-2006 15:04:05")}

	g.Header("Content-Type", Constant.ContentTypeJSON)
	g.JSON(http.StatusOK, responseData)
}

func HttpResponseError(g *gin.Context, data interface{}, code int, err error) {
	location, _ := time.LoadLocation(Constant.TimeLocation)
	responseData := setResponse{
		Status:     err.Error(),
		AccessTime: time.Now().In(location).Format("02-01-2006 15:04:05"),
		Data:       data,
	}
	g.Header("Content-Type", Constant.ContentTypeJSON)
	g.JSON(code, responseData)
}

func SetResponsePublisher(status string, data interface{}, code int) ResponsePublisher {
	responseData := ResponsePublisher{
		Status: status,
		Data:   data,
		Code:   code,
	}
	return responseData
}

func GetResponseFromRPCReply(message []byte) (response ResponsePublisher) {
	json.Unmarshal(message, &response)
	return
}
