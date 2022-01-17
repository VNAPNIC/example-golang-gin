package dto

type (
	User struct {
		ID       uint   `json:"id"`
		UserName string `json:"user_name"`
		Status   int    `json:"status"`
		Role     Role   `json:"role"`
	}

	Auth struct {
		Username string `json:"user_name" validate:"required,min=4,max=20" minLength:"4" maxLength:"20"`
		Password string `json:"password" validate:"required,min=4,max=20" minLength:"4" maxLength:"20"`
	}

	AddUser struct {
		Auth
		RoleId uint `json:"role_id" validate:"omitempty,numeric,min=0"`
	}
)
