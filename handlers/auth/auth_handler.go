package authHandler

import (
	"net/http"
	"serverhealthcarepanel/dto"
	userService "serverhealthcarepanel/services/user"
	"serverhealthcarepanel/utils"
	"serverhealthcarepanel/utils/code"
	"serverhealthcarepanel/utils/response"

	"github.com/labstack/echo/v4"
)

// UserLogin
// @Summary User Login
// @Description User Login
// @Accept json
// @Produce json
// @Tags Auth
// @Param payload body dto.Auth true "user login"
// @Success 200 {object} response.Struct
// @Failure 400 {object} response.Struct "wrong request parameter"
// @Failure 401 {object} response.Struct "The corresponding username or password is incorrect"
// @Router /login [post]
func UserLogin(ctx echo.Context) error {
	auth := new(dto.Auth)

	if err := ctx.Bind(&auth); err != nil {
		return response.Error(ctx, http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
	}

	if err := ctx.Validate(*auth); err != nil {
		return response.Error(ctx, http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
	}

	RError, isExist, user := userService.CheckAuth(auth)

	RCode := code.ErrorUserPasswordInvalid

	if isExist {
		token, err := utils.GenerateToken(utils.Claims{
			UserId:   user.ID,
			Username: user.Username,
			RoleKey:  user.Role.RoleKey,
			IsAdmin:  user.Role.IsAdmin,
		})

		if err != nil {
			RError = err
			RCode = code.ErrorAuthToken
		} else {
			// set login time
			userService.SetLoggedTime(user.ID)
			return response.Success(ctx, map[string]string{"token": token})
		}

	}

	return response.Error(
		ctx,
		http.StatusUnauthorized,
		RCode,
		code.GetMsg(RCode),
		utils.DetectError(RError),
	)
}
