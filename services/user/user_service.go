package userService

import (
	"log"
	"serverhealthcarepanel/dto"
	"serverhealthcarepanel/models"
	"serverhealthcarepanel/utils"
	"strings"
	"time"
)

func CheckAuth(auth *dto.Auth) (interface{}, bool, model.Auth) {
	return model.CheckAuth(auth.Username, auth.Password)
}

// SetLoggedTime
// Set login time
func SetLoggedTime(userId uint) {
	wheres := make(map[string]interface{})
	wheres["id"] = userId

	updates := make(map[string]interface{})
	updates["logged_in_at"] = time.Now()
	_, rowAffected := model.Update(&model.Auth{}, wheres, updates)
	if rowAffected == 0 {
		log.Println("Failed to set login time!")
	}
}

func CreateUser(newUser *dto.AddUser) error {
	return model.CreateUser(model.Auth{
		Username: strings.TrimSpace(newUser.Username),
		Password: utils.EncodeUserPassword(newUser.Password),
		RoleId:   newUser.RoleId,
	})
}

func ChangePassword(userId uint, password string) (bool, model.Auth) {
	return model.ChangePassword(userId, utils.EncodeUserPassword(password))
}
