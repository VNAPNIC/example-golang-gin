package user

import (
	model "serverhealthcarepanel/models"
	"serverhealthcarepanel/utils"
	"strings"
)

type (
	AuthStruct struct {
		Username string `json:"user_name" validate:"required,min=4,max=20" minLength:"4" maxLength:"20"`
		Password string `json:"password" validate:"required,min=4,max=20" minLength:"4" maxLength:"20"`
	}

	AddUserStruct struct {
		AuthStruct
		RoleId int `json:"role_id" validate:"omitempty,numeric,min=0"`
	}
)

func CreateUser(newUser *AddUserStruct) error {
	return model.CreateUser(model.Auth{
		Username: strings.TrimSpace(newUser.Username),
		Password: utils.EncodeUserPassword(newUser.Password),
		RoleId:   newUser.RoleId,
	})
}
