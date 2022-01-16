package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"serverhealthcarepanel/utils/code"
	"time"
)

// JSONTime format json time field by myself
type JSONTime struct {
	time.Time
}

type Struct struct {
	TimeStamp JSONTime    `json:"time"`
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

func Error(ctx echo.Context, httpCode, errorCode int, msg string, data interface{}) error {
	return ctx.JSON(httpCode, Struct{
		Code:    errorCode,
		Message: msg,
		Data:    data,
	})
}

func Success(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusOK, Struct{
		Code:    code.SUCCESS,
		Message: code.GetMsg(code.SUCCESS),
		Data:    data,
	})
}
