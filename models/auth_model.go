package models

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Active   bool   `json:"actived"`
	RoleId   bool   `json:"role_id"`
}
