package authHandler

import (
	"net/http"
	"serverhealthcarepanel/dto"
	model "serverhealthcarepanel/models"
	userService "serverhealthcarepanel/services/user"
	"serverhealthcarepanel/utils"
	"serverhealthcarepanel/utils/code"
	jwtUtil "serverhealthcarepanel/utils/jwt"
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
		token, err := jwtUtil.GenerateToken(jwtUtil.Claims{
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
		http.StatusOK,
		RCode,
		code.GetMsg(RCode),
		utils.DetectError(RError.(error)),
	)
}

// @Summary User Logout
// @Description 用户登出
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Tags User
// @Success 200 {object} common.Response
// @Router /user/logout [put]
// func UserLogout(ctx echo.Context) {

// 	user := jwtUtil.GetClaim(ctx)

// 	userService.JoinBlockList(user.UserId, c.GetHeader("Authorization")[7:])

// 	appG.Response(http.StatusOK, code.SUCCESS, "ok", nil)
// }

// ChangePassword
// @Summary Change password
// @Description Change password
// @Accept json
// @Produce json
// @Tags User
// @Param payload body dto.ChangePassword true "user change password"
// @Success 200 {object} response.Struct
// @Router /user/change_password [put]
func ChangePassword(ctx echo.Context) error {
	password := new(dto.ChangePassword)

	if err := ctx.Bind(&password); err != nil {
		return response.Error(ctx, http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
	}

	if err := ctx.Validate(*password); err != nil {
		return response.Error(ctx, http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
	}

	user := jwtUtil.GetClaim(ctx)

	_, isExist, auth := model.CheckAuth(user.Username, password.OldPassword)
	if !isExist {
		return response.Error(ctx, http.StatusOK, code.ErrorUserOldPasswordInvalid, code.GetMsg(code.ErrorUserOldPasswordInvalid), nil)
	}

	if successful, _ := userService.ChangePassword(auth.ID, password.NewPassword); !successful {
		return response.Error(ctx, http.StatusOK, code.UnknownError, code.GetMsg(code.UnknownError), nil)
	}

	return response.Success(ctx, nil)
}
