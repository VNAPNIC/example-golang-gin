package authHalder

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"serverhealthcarepanel/dto"
	"serverhealthcarepanel/services/user"
	"serverhealthcarepanel/utils/code"
	"serverhealthcarepanel/utils/response"
)

// UserLogin
// @Summary User Login
// @Description User Login
// @Accept json
// @Produce json
// @Tags Auth
// @Param payload body services.AuthStruct true "user login"
// @Success 200 {object} response.Struct
// @Failure 400 {object} response.Struct "wrong request parameter"
// @Failure 401 {object} response.Struct "The corresponding username or password is incorrect"
// @Router /login [post]
func UserLogin(ctx echo.Context) error {
	userLogin := new(user.AuthStruct)

	if err := ctx.Bind(&userLogin); err != nil {
		return response.Error(ctx, http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
	}

	if err := ctx.Validate(&userLogin); err != nil {
		return response.Error(ctx, http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
	}

	if userLogin.Username != "hainam" || userLogin.Password != "hainam" {
		return response.Error(ctx, http.StatusUnauthorized, code.ErrorUserPasswordInvalid, code.GetMsg(code.ErrorUserPasswordInvalid), nil)
	}

	return response.Success(ctx, dto.UserDto{
		Name:  "Hai nam",
		Token: "aaaaaaaaaaaaaa",
	})
}
