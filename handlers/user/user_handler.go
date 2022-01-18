package userHandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"serverhealthcarepanel/dto"
	"serverhealthcarepanel/services/user"
	"serverhealthcarepanel/utils/code"
)

// CreateUser
// @Summary create user
// @Description Create new user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Tags User
// @Param payload body dto.AddUser true "create new user"
// @Success 200 {object} dto.Struct
// @Failure 400 {object} dto.Struct "wrong request parameter"
// @Failure 500 {object} dto.Struct
// @Router /user [post]
func CreateUser(ctx echo.Context) error {
	newUser := new(dto.AddUser)

	if err := ctx.Bind(&newUser); err != nil {
		return dto.Error(ctx, http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
	}

	if err := ctx.Validate(*newUser); err != nil {
		return dto.Error(ctx, http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
	}

	if err := userService.CreateUser(newUser); err != nil {
		return dto.Error(ctx, http.StatusOK, code.ErrorFailedAddNewUser, code.GetMsg(code.ErrorFailedAddNewUser), err.Error())
	}

	return dto.Success(ctx, nil)
}
