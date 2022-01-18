package authHandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"serverhealthcarepanel/dto"
	"serverhealthcarepanel/models"
	"serverhealthcarepanel/services/user"
	"serverhealthcarepanel/utils"
	"serverhealthcarepanel/utils/code"
	"serverhealthcarepanel/utils/jwt"
	"serverhealthcarepanel/utils/redis"
)

// UserLogin
// @Summary User Login
// @Description User Login
// @Accept json
// @Produce json
// @Tags Auth
// @Param payload body dto.Auth true "user login"
// @Success 200 {object} dto.Struct
// @Failure 400 {object} dto.Struct "wrong request parameter"
// @Failure 401 {object} dto.Struct "The corresponding username or password is incorrect"
// @Router /login [post]
func UserLogin(ctx echo.Context) error {
	auth := new(dto.Auth)

	if err := ctx.Bind(&auth); err != nil {
		return dto.Error(ctx, http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
	}

	if err := ctx.Validate(*auth); err != nil {
		return dto.Error(ctx, http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
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
			RError = err.Error()
			RCode = code.ErrorAuthToken
		} else {
			// set login time
			userService.SetLoggedTime(user.ID)
			return dto.Success(ctx, map[string]string{"token": token})
		}

	}

	return dto.Error(
		ctx,
		http.StatusOK,
		RCode,
		code.GetMsg(RCode),
		utils.DetectError(RError.(error)),
	)
}

// UserLogout
// @Summary User Logout
// @Description User logout
// @Accept json
// @Produce json
// @Tags User
// @Success 200 {object} dto.Struct
// @Router /user/logout [put]
func UserLogout(ctx echo.Context) error {
	user := jwtUtil.GetClaim(ctx)

	successful, err := redisUtil.Delete(user.Issuer)
	if err != nil {
		return dto.Error(ctx, http.StatusOK, code.UnknownError, code.GetMsg(code.UnknownError), err.Error())
	}

	if successful == false {
		return dto.Error(ctx, http.StatusOK, code.UnknownError, code.GetMsg(code.UnknownError), nil)
	}

	return dto.Success(ctx, nil)
}

// ChangePassword
// @Summary Change password
// @Description Change password
// @Accept json
// @Produce json
// @Tags User
// @Param payload body dto.ChangePassword true "user change password"
// @Success 200 {object} dto.Struct
// @Router /user/change_password [put]
func ChangePassword(ctx echo.Context) error {
	password := new(dto.ChangePassword)

	if err := ctx.Bind(&password); err != nil {
		return dto.Error(ctx, http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
	}

	if err := ctx.Validate(*password); err != nil {
		return dto.Error(ctx, http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
	}

	user := jwtUtil.GetClaim(ctx)

	_, isExist, auth := model.CheckAuth(user.Username, password.OldPassword)
	if !isExist {
		return dto.Error(ctx, http.StatusOK, code.ErrorUserOldPasswordInvalid, code.GetMsg(code.ErrorUserOldPasswordInvalid), nil)
	}

	if successful, _ := userService.ChangePassword(auth.ID, password.NewPassword); !successful {
		return dto.Error(ctx, http.StatusOK, code.UnknownError, code.GetMsg(code.UnknownError), nil)
	}

	return dto.Success(ctx, nil)
}
