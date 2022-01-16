package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
	mMiddleware "serverhealthcarepanel/middleware"
	model "serverhealthcarepanel/models"
	"serverhealthcarepanel/routers"
	"serverhealthcarepanel/utils"
	"serverhealthcarepanel/utils/setting"
)

func init() {
	setting.Setup()
	model.Setup()
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
