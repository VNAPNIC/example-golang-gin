package services

type (
	AuthStruct struct {
		Username string `json:"user_name" validate:"required,min=6,max=20" minLength:"6" maxLength:"20"`
		Password string `json:"password" validate:"required,min=6,max=20" minLength:"6" maxLength:"20"`
	}
)
