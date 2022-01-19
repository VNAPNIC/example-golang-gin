package userHandler

import (
	"github.com/gin-gonic/gin"
	"healthcare-panel/common"
	"healthcare-panel/dto"
	userService "healthcare-panel/services/user"
	"healthcare-panel/utils/code"
	"net/http"
)

// @Summary create user
// @Description Create new user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Tags User
// @Param payload body dto.AddUser true "create new user"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response "wrong request parameter"
// @Failure 500 {object} common.Response
// @Router /user [post]
func CreateUser(ctx *gin.Context) {
	g := common.Gin{C: ctx}

	newUser := new(dto.AddUser)

	if err := ctx.Bind(&newUser); err != nil {
		g.Error(http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
		return
	}

	if err := common.CheckBindStructParameter(*newUser); err != nil {
		g.Error(http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
		return
	}

	if err := userService.CreateUser(newUser); err != nil {
		g.Error(http.StatusOK, code.ErrorFailedAddNewUser, code.GetMsg(code.ErrorFailedAddNewUser), err.Error())
		return
	}

	g.Success(nil)
}
