package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"serverhealthcarepanel/dto"
	"serverhealthcarepanel/services"
	"serverhealthcarepanel/utils/code"
	"serverhealthcarepanel/utils/response"
)

// UserLogin
// @Summary User Login
// @Description User login
// @Accept json
// @Produce json
// @Tags Auth
// @Param payload body services.AuthStruct true "user login"
// @Success 200 {object}
// @Router /login [post]
func UserLogin(ctx echo.Context) error {
	userLogin := new(services.AuthStruct)

	if err := ctx.Bind(userLogin); err != nil {
		return response.Response(ctx, http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
	}

	if err := ctx.Validate(userLogin); err != nil {
		return response.Response(ctx, http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
	}

	if userLogin.Username != "hainam" || userLogin.Password != "hainam" {
		return response.Response(ctx, http.StatusUnauthorized, code.ErrorUserPasswordInvalid, code.GetMsg(code.ErrorUserPasswordInvalid), nil)
	}

	return response.Success(ctx, dto.UserDto{
		Name:  "Hai nam",
		Token: "aaaaaaaaaaaaaa",
	})
}
