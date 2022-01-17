package userService

import (
	"serverhealthcarepanel/dto"
	"serverhealthcarepanel/models"
	"serverhealthcarepanel/utils"
	"strings"
)

func CheckAuth(auth *dto.Auth) (error, bool, model.Auth) {
	return model.CheckAuth(auth.Username, auth.Password)
}

func SetLoggedTime(userId uint) {

}

func CreateUser(newUser *dto.AddUser) error {
	return model.CreateUser(model.Auth{
		Username: strings.TrimSpace(newUser.Username),
		Password: utils.EncodeUserPassword(newUser.Password),
		RoleId:   newUser.RoleId,
	})
}
