package dto

type RoleDto struct {
	RoleId   int    `json:"role_id"`
	RoleName string `json:"role_name"`
	IsAdmin  bool   `json:"is_admin"`
	RoleKey  string `json:"role_key"`
}
