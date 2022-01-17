package role

import (
	"serverhealthcarepanel/dto"
	"serverhealthcarepanel/models"
)

func CreateRole(newRole *dto.CreateRole) error {
	return model.CreateRole(model.Role{
		RoleKey:  newRole.RoleKey,
		RoleName: newRole.RoleName,
	})

}
