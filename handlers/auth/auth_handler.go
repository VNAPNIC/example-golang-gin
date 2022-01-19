package authHandler

import (
	"github.com/gin-gonic/gin"
	"healthcare-panel/common"
	"healthcare-panel/dto"
	model "healthcare-panel/models"
	userService "healthcare-panel/services/user"
	"healthcare-panel/utils"
	"healthcare-panel/utils/code"
	jwtUtil "healthcare-panel/utils/jwt"
	redisUtil "healthcare-panel/utils/redis"
	"net/http"
)

// UserLogin
// @Summary User Login
// @Description User Login
// @Accept json
// @Produce json
// @Tags Auth
// @Param payload body dto.Auth true "user login"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response "wrong request parameter"
// @Failure 401 {object} common.Response "The corresponding username or password is incorrect"
// @Router /login [post]
func UserLogin(ctx *gin.Context) {
	g := common.Gin{C: ctx}

	auth := new(dto.Auth)

	if err := g.C.Bind(&auth); err != nil {
		g.Error(http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
		return
	}

	if err := common.CheckBindStructParameter(*auth); err != nil {
		g.Error(http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
		return
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
			g.Success(map[string]string{"token": token})
			return
		}

	}

	g.Error(
		http.StatusOK,
		RCode,
		code.GetMsg(RCode),
		utils.DetectError(RError.(error)),
	)
}

// @Summary User Logout
// @Description User logout
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Tags User
// @Success 200 {object} common.Response
// @Router /user/logout [put]
func UserLogout(ctx *gin.Context) {
	g := common.Gin{C: ctx}

	user, _ := jwtUtil.GetClaim(g.C)

	successful, err := redisUtil.Delete(user.Issuer)
	if err != nil {
		g.Error(http.StatusOK, code.UnknownError, code.GetMsg(code.UnknownError), err.Error())
		return
	}

	if successful == false {
		g.Error(http.StatusOK, code.UnknownError, code.GetMsg(code.UnknownError), nil)
		return
	}

	g.Success(nil)
}

// @Summary Change password
// @Description Change password
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Tags User
// @Param payload body dto.ChangePassword true "user change password"
// @Success 200 {object} common.Response
// @Router /user/change_password [put]
func ChangePassword(ctx *gin.Context) {
	g := common.Gin{C: ctx}

	password := new(dto.ChangePassword)

	if err := ctx.Bind(&password); err != nil {
		g.Error(http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
		return
	}

	if err := common.CheckBindStructParameter(*password); err != nil {
		g.Error(http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
		return
	}

	user, _ := jwtUtil.GetClaim(g.C)

	_, isExist, auth := model.CheckAuth(user.Username, password.OldPassword)
	if !isExist {
		g.Error(http.StatusOK, code.ErrorUserOldPasswordInvalid, code.GetMsg(code.ErrorUserOldPasswordInvalid), nil)
		return
	}

	if successful, _ := userService.ChangePassword(auth.ID, password.NewPassword); !successful {
		g.Error(http.StatusOK, code.UnknownError, code.GetMsg(code.UnknownError), nil)
		return
	}

	g.Success(nil)
}
