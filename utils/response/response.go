package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"serverhealthcarepanel/utils/code"
	"time"
)

type Struct struct {
	TimeStamp int64       `json:"time_stamp"`
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

func (e Struct) AddTimeStamp() Struct {
	e.TimeStamp = time.Now().UnixMilli()
	return e
}

func Response(ctx echo.Context, httpCode, errorCode int, msg string, data interface{}) error {
	return ctx.JSON(httpCode, Struct{
		TimeStamp: time.Now().UnixMilli(),
		Code:      errorCode,
		Message:   msg,
		Data:      data,
	})
}

func Success(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusOK, Struct{
		TimeStamp: time.Now().UnixMilli(),
		Code:      code.SUCCESS,
		Message:   code.GetMsg(code.SUCCESS),
		Data:      data,
	})
}
