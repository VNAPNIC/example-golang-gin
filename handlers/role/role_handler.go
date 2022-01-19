package roleHandler

import (
	"github.com/gin-gonic/gin"
	"healthcare-panel/common"
	"healthcare-panel/dto"
	"healthcare-panel/services/role"
	"healthcare-panel/utils/code"
	"net/http"
)

// @Summary Create role
// @Description Create role
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Tags Role
// @Param role_id path int true "role_id"
// @Param payload body dto.CreateRole true "YES"、
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response "wrong request parameter"
// @Router /role [post]
func CreateRole(ctx *gin.Context) {
	g := common.Gin{C: ctx}

	newRole := new(dto.CreateRole)

	if err := ctx.Bind(&newRole); err != nil {
		g.Error(http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
		return
	}

	if err := common.CheckBindStructParameter(*newRole); err != nil {
		g.Error(http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
		return
	}

	if err := role.CreateRole(newRole); err != nil {
		g.Error(http.StatusOK, code.ErrorFailedAddNewRole, code.GetMsg(code.ErrorFailedAddNewRole), err.Error())
		return
	}

	g.Success(nil)
}
