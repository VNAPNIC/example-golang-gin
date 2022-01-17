package dto

type (
	Role struct {
		RoleId   uint   `json:"role_id"`
		RoleName string `json:"role_name"`
		IsAdmin  bool   `json:"is_admin"`
		RoleKey  string `json:"role_key"`
	}

	CreateRole struct {
		RoleName string `json:"role_name"`
		IsAdmin  bool   `json:"is_admin"`
		RoleKey  string `json:"role_key"`
	}

	UpdateRole struct {
		RoleName string `json:"role_name"`
		IsAdmin  bool   `json:"is_admin"`
		RoleKey  string `json:"role_key"`
	}
)
