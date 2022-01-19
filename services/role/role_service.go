package role

import (
	"healthcare-panel/dto"
	model "healthcare-panel/models"
)

func CreateRole(newRole *dto.CreateRole) error {
	return model.CreateRole(model.Role{
		RoleKey:  newRole.RoleKey,
		RoleName: newRole.RoleName,
	})

}
