package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	mMiddleware "serverhealthcarepanel/middleware"
	"serverhealthcarepanel/routers"
	"serverhealthcarepanel/utils"
	"serverhealthcarepanel/utils/setting"
)

func init() {
	setting.Setup()
	//TODO something
}

func main() {
	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	e.Use(middleware.Logger(), middleware.Recover(), mMiddleware.CORS())
	routers.InitRouter(e)
	e.Logger.Fatal(e.Start(":1323"))
}
