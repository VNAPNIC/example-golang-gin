package dto

type UserDto struct {
	ID       int     `json:"id"`
	UserName string  `json:"user_name"`
	Status   int     `json:"status"`
	Role     RoleDto `json:"role"`
	Token    string  `json:"token"`
}
