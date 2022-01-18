package roleHandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"serverhealthcarepanel/dto"
	"serverhealthcarepanel/services/role"
	"serverhealthcarepanel/utils/code"
)

// CreateRole
// @Summary Create role
// @Description Create role
// @Accept json
// @Produce json
// @Tags Role
// @Param role_id path int true "role_id"
// @Param payload body dto.CreateRole true "YES"、
// @Success 200 {object} dto.Struct
// @Failure 400 {object} dto.Struct "wrong request parameter"
// @Router /role [post]
func CreateRole(ctx echo.Context) error {
	newRole := new(dto.CreateRole)

	if err := ctx.Bind(&newRole); err != nil {
		return dto.Error(ctx, http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
	}

	if err := ctx.Validate(*newRole); err != nil {
		return dto.Error(ctx, http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
	}

	if err := role.CreateRole(newRole); err != nil {
		return dto.Error(ctx, http.StatusOK, code.ErrorFailedAddNewRole, code.GetMsg(code.ErrorFailedAddNewRole), err.Error())
	}

	return dto.Success(ctx, nil)
}
