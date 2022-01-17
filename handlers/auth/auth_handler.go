package authHalder

import (
	"net/http"
	"serverhealthcarepanel/dto"
	userService "serverhealthcarepanel/services/user"
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
// @Param payload body services.AuthStruct true "user login"
// @Success 200 {object} response.Struct
// @Failure 400 {object} response.Struct "wrong request parameter"
// @Failure 401 {object} response.Struct "The corresponding username or password is incorrect"
// @Router /login [post]
func UserLogin(ctx echo.Context) error {
	auth := new(userService.AuthStruct)

	if err := ctx.Bind(&auth); err != nil {
		return response.Error(ctx, http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
	}

	if err := ctx.Validate(*auth); err != nil {
		return response.Error(ctx, http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
	}

	var err, isExist, user = userService.CheckAuth(auth)

	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, code.ErrorUserPasswordInvalid, code.GetMsg(code.ErrorUserPasswordInvalid), err.Error())
	}

	if !isExist {
		return response.Error(ctx, http.StatusUnauthorized, code.ErrorUserPasswordInvalid, code.GetMsg(code.ErrorUserPasswordInvalid), nil)
	}

	return response.Success(ctx, dto.UserDto{
		ID:       user.ID,
		UserName: user.Username,
		Status:   user.Status,
		Role: dto.RoleDto{
			RoleId:   user.Role.RoleId,
			RoleName: user.Role.RoleName,
			IsAdmin:  user.Role.IsAdmin,
			RoleKey:  user.Role.RoleKey,
		},
		Token: "aaaaaaaaaaasssssssaaa",
	})
}
