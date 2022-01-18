package main

import (
	mMiddleware "serverhealthcarepanel/middleware"
	model "serverhealthcarepanel/models"
	"serverhealthcarepanel/routers"
	"serverhealthcarepanel/utils"
	redisUtil "serverhealthcarepanel/utils/redis"
	"serverhealthcarepanel/utils/setting"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

func init() {
	setting.Setup()
	model.Setup()
	redisUtil.Setup()
}

// @title Healthcare panel
// @version 1.0
// @in header like: Bearer xxxx
// @name Authorization
func main() {
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	e.Use(middleware.Logger(), middleware.Recover(), mMiddleware.CORS())
	routers.InitRouter(e)
	e.Logger.Fatal(e.Start(":1323"))
}
