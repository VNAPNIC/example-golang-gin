package common

import (
	"github.com/gin-gonic/gin"
	"healthcare-panel/utils/code"
	"net/http"
	"time"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	String  time.Time   `json:"time"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (g *Gin) Error(httpCode, errorCode int, msg string, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code:    errorCode,
		Message: msg,
		Data:    data,
	})
}

func (g *Gin) Success(data interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code:    code.SUCCESS,
		Message: code.GetMsg(code.SUCCESS),
		Data:    data,
	})
}
