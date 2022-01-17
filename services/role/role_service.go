package role

import model "serverhealthcarepanel/models"

type (
	RoleStruct struct {
		PageNum  int
		PageSize int
	}

	NewRoleStruct struct {
		RoleKey string `json:"role_key" form:"role_key" validate:"required,min=4,max=10" minLength:"4" maxLength:"10"`
	}

	UpdateRoleStruct struct {
		// Role name
		RoleName string `json:"role_name" form:"role_name" validate:"required,min=4,max=10" minLength:"4" maxLength:"10"`
	}

	CreateRoleStruct struct {
		NewRoleStruct
		UpdateRoleStruct
	}
)

func CreateRole(newRole *CreateRoleStruct) error {
	return model.CreateRole(model.Role{
		RoleKey:  newRole.RoleKey,
		RoleName: newRole.RoleName,
	})

}
