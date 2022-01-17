package model

import (
	_ "gorm.io/gorm"
)

type Auth struct {
	BaseModel
	RoleId     int      `gorm:"DEFAULT:0;NOT NULL;" json:"role_id"`
	Status     int      `gorm:"type:int(1);DEFAULT:0;NOT NULL;" json:"status"`
	Username   string   `gorm:"Size:20;uniqueIndex;NOT NULL;" json:"user_name"`
	Password   string   `gorm:"Size:50;NOT NULL;" json:"-"`
	RoleName   string   `gorm:"-" json:"role_name"`
	LoggedInAt JSONTime `gorm:"type:timestamp;default:current_timestamp" json:"logged_in_at"`
	Role       Role     `gorm:"references:RoleId"`
}

func (Auth) TableName() string {
	return TablePrefix + "auth"
}

func CreateUser(auth Auth) error {
	res := db.Create(&auth)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}
